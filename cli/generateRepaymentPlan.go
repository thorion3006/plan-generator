package cli

import (
	"fmt"
	"strings"

	plangenerator "github.com/thorion3006/plan-generator"
)

// Calls RepaymentPlan, printTable and appendFile functions as needed.
func generateRepaymentPlan(loan *plangenerator.LoanDetails) {
	repaymentPlan, err := plangenerator.RepaymentPlan(loan)
	if err != nil {
		fmt.Println(err)
		return
	}
	table := printTable(repaymentPlan)

	fmt.Println()
	fmt.Println("Repayment Plan:")
	for _, row := range table {
		fmt.Print(row)
	}
	fmt.Print("Do you want to save this to the file?(y/N):")
	userResponse := UserResponseProcessor()

	userResponse = strings.ToLower(userResponse)
	if userResponse == "y" || userResponse == "yes" {
		appendFile(table, loan)
	}
}
