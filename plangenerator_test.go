package plangenerator

import (
	"fmt"
	"testing"
	"time"
)

func TestLoanDetailsValidation(t *testing.T) {
	// arrange
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	loan := LoanDetails{LoanAmount: 0, Duration: 24, NominalRate: 5, StartDate: date}

	// act
	err := loan.Validate()

	// assert
	if err == nil {
		t.Fatal("Expected validation to produce an error, but got none.")
	}
}

func TestAnnuityCalculation(t *testing.T) {
	// arrange
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	loan := LoanDetails{LoanAmount: 5000, Duration: 24, NominalRate: 5, StartDate: date}
	annuity := 219.35694867034132

	// act
	loan.AnnuityCalculation()

	// assert
	if loan.Annuity != annuity {
		t.Fatalf("Expected annuity to be equal to %v but got %v.", annuity, loan.Annuity)
	}
}

func TestRepaymentPlan(t *testing.T) {
	// arrange
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	date2, _ := time.Parse("02-01-2006", "01-12-2019")
	loan := LoanDetails{LoanAmount: 5000, Duration: 24, NominalRate: 5, StartDate: date}
	repayment1 := MonthlyRepayment{
		BorrowerPaymentAmount: 219.35694867034132,
		Date: date,
		InitialOutstandingPrincipal:   5000,
		Interest:                      20.833333333333332,
		Principal:                     198.52361533700798,
		RemainingOutstandingPrincipal: 4801.476384662992,
	}
	repayment2 := MonthlyRepayment{
		BorrowerPaymentAmount: 219.35694867034132,
		Date: date2,
		InitialOutstandingPrincipal:   218.44675386260155,
		Interest:                      0.9101948077608397,
		Principal:                     218.4467538625805,
		RemainingOutstandingPrincipal: 0,
	}

	// act
	result, _ := RepaymentPlan(&loan)

	// assert
	if len(result) != 24 {
		t.Errorf("Expected result to be of length 24 but got %v.", len(result))
	}

	if result[0] != repayment1 || result[23] != repayment2 {
		t.Error("Expected result not equal to actual repayment plan.")
	}
}

func ExampleRepaymentPlan() {
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	loan := LoanDetails{LoanAmount: 1000, NominalRate: 12, Duration: 2, StartDate: date}
	repaymentPlan, err := RepaymentPlan(&loan)
	if err != nil {
		return
	}
	fmt.Println(repaymentPlan)

	// Output:
	// [{507.5124378109438 2018-01-01 00:00:00 +0000 UTC 1000 10 497.5124378109438 502.4875621890562} {507.5124378109438 2018-02-01 00:00:00 +0000 UTC 502.4875621890562 5.024875621890562 502.4875621890533 0}]
}
