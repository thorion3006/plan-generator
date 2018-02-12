package cli

import (
	"fmt"
	"log"
	"strconv"
	"time"

	plangenerator "github.com/thorion3006/plan-generator"
)

// Assigns user input values to the LoanDetails struct except for the Annuity field.
func inputLoanDetails(loan *plangenerator.LoanDetails) {
	var err error
	fmt.Println("Welcome to Loan repayment Calculator!")

	fmt.Println("Please input the loan details:")
	//Input for loan amount
	fmt.Print("Loan amount (format: 1234.56): ")
	userResponse := UserResponseProcessor()
	loan.LoanAmount, err = strconv.ParseFloat(userResponse, 64)
	if err != nil {
		log.Fatal(err)
	}

	//Input for loan duration
	fmt.Print("Loan Duration in months (format: 12): ")
	userResponse = UserResponseProcessor()
	var duration int64
	duration, err = strconv.ParseInt(userResponse, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	loan.Duration = int(duration)

	//Input for nominal interest rate
	fmt.Print("Nominal Interest Rate (format: 12.56): ")
	userResponse = UserResponseProcessor()
	loan.NominalRate, err = strconv.ParseFloat(userResponse, 64)
	if err != nil {
		log.Fatal(err)
	}

	//Input for repayment start date
	fmt.Print("Repayment start date (format: 31-01-2006): ")
	userResponse = UserResponseProcessor()
	loan.StartDate, err = time.Parse("02-01-2006", userResponse)
	if err != nil {
		log.Fatal(err)
	}
}
