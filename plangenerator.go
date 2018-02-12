package plangenerator

import (
	"fmt"
	"math"
	"time"
)

// Struct to store the input loan details for repayment plan generation.
// Note: Duration field takes the duration of the loan period in months and not years.
type LoanDetails struct {
	LoanAmount  float64   `json:"loanAmount"`
	NominalRate float64   `json:"nominalRate"`
	Duration    int       `json:"duration"`
	StartDate   time.Time `json:"startDate"`
	Annuity     float64   `json:"annuity"`
}

// Struct to store the repayment details for each month of the loan period.
type MonthlyRepayment struct {
	BorrowerPaymentAmount         float64   `json:"borrowerPaymentAmount"`
	Date                          time.Time `json:"date"`
	InitialOutstandingPrincipal   float64   `json:"initialOutstandingPrincipal"`
	Interest                      float64   `json:"interest"`
	Principal                     float64   `json:"principal"`
	RemainingOutstandingPrincipal float64   `json:"remainingOutstandingPrincipal"`
}

// Validates if the LoanAmount, NominalRate and Duration are greater than 1 and StartDate is after 01-01-2000.
func (loan *LoanDetails) Validate() (err error) {
	var errs, err1, err2, err3, err4 string
	jan2000, _ := time.Parse("02-01-2006", "01-01-2000")
	if loan.LoanAmount < 1 {
		err1 = "Loan Amount can not be zero or less. "
		errs += err1
	}
	if loan.Duration < 1 {
		err2 = "Duration can not be zero or less. "
		errs += err2
	}
	if loan.NominalRate < 1 {
		err3 = "Nominal Interest Rate can not be negative. "
		errs += err3
	}
	if loan.StartDate.Before(jan2000) {
		err4 = "Start date can not be before 01-01-2000."
		errs += err4
	}
	if len(errs) > 0 {
		err = fmt.Errorf("Error: %s", errs)
	}
	return
}

// Populates the Annuity field when called invoked.
func (loan *LoanDetails) AnnuityCalculation() {
	loan.Annuity = (loan.LoanAmount * (loan.NominalRate / 1200)) / (1 - (math.Pow(1+(loan.NominalRate/1200), float64(-loan.Duration))))
}

// Calculates the repayment plan for an annuity loan.
// It takes a reference to a LoanDetails struct and returns a slice of MonthlyRepayment struct.
// The LoanAmount, NominalRate, Duration and StartDate fields of the LoanDetails struct are mandatory.
// It invokes the Validate and AnnuityCalcuation functions of the LoanDetails struct to validate the input and calculate annuity.
func RepaymentPlan(loan *LoanDetails) (repaymentPlan []MonthlyRepayment, err error) {
	err = loan.Validate()
	if err != nil {
		return
	}
	loan.AnnuityCalculation()
	intialOutstandingPrincipal := loan.LoanAmount
	repaymentPlan = make([]MonthlyRepayment, loan.Duration)
	for month := 0; month < loan.Duration; month++ {
		repaymentPlan[month].Date = loan.StartDate.AddDate(0, month, 0)
		repaymentPlan[month].InitialOutstandingPrincipal = intialOutstandingPrincipal
		repaymentPlan[month].Interest = (loan.NominalRate * 30 * intialOutstandingPrincipal) / 36000
		repaymentPlan[month].Principal = loan.Annuity - repaymentPlan[month].Interest
		repaymentPlan[month].BorrowerPaymentAmount = repaymentPlan[month].Interest + repaymentPlan[month].Principal
		repaymentPlan[month].RemainingOutstandingPrincipal = intialOutstandingPrincipal - repaymentPlan[month].Principal
		if repaymentPlan[month].RemainingOutstandingPrincipal < 1 {
			repaymentPlan[month].RemainingOutstandingPrincipal = 0
		}
		intialOutstandingPrincipal = repaymentPlan[month].RemainingOutstandingPrincipal
	}
	return
}
