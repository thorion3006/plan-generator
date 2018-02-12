package cli

import (
	"fmt"
	"strings"

	"github.com/thorion3006/plan-generator"
)

// Starts the CLI application for the plangenerator.
func Start() {
	var loan plangenerator.LoanDetails
	for {
		inputLoanDetails(&loan)
		generateRepaymentPlan(&loan)
		fmt.Println()
		fmt.Print("Do you want to calculate a new loan payment?(y/N): ")
		userResponse := UserResponseProcessor()
		userResponse = strings.ToLower(userResponse)
		if userResponse == "y" || userResponse == "yes" {
			ClearScreen()
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Println("Thank you for using plan-generator!")
}
