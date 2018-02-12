package cli

import (
	"fmt"

	plangenerator "github.com/thorion3006/plan-generator"
)

// Formats the MonthlyRepayment slice returned by the RepaymentPlan function into a tabular form stored in a slice of strings.
func printTable(repaymentPlan []plangenerator.MonthlyRepayment) (table []string) {
	table = make([]string, 4+(2*len(repaymentPlan)))
	table[0] = fmt.Sprint("----------------------------------------------------------------------------------------------------\n")
	table[1] = fmt.Sprintf("|%3v|%10v|%18v|%10v|%10v|%20v|%21v|\n", "#", "Date", "Annuity (Borrower", "Principal", "Interest", "Initial Outstanding", "Remaining Outstanding")
	table[2] = fmt.Sprintf("|%3v|%10v|%18v|%10v|%10v|%20v|%21v|\n", "", "", "Payment Amount)", "", "", "Principal", "Principal")
	table[3] = fmt.Sprintf("|%3v|%10v|%18v|%10v|%10v|%20v|%21v|\n", "---", "----------", "------------------", "----------", "----------", "--------------------", "---------------------")
	tableIndex := 4
	for index := 0; index < len(repaymentPlan); index++ {
		table[tableIndex] = fmt.Sprintf("|%3d|%10v|%17.2f%s|%9.2f%s|%9.2f%s|%19.2f%s|%20.2f%s|\n",
			index+1,
			repaymentPlan[index].Date.Format("02-01-2006"),
			repaymentPlan[index].BorrowerPaymentAmount,
			"\u20AC",
			repaymentPlan[index].Principal,
			"\u20AC",
			repaymentPlan[index].Interest,
			"\u20AC",
			repaymentPlan[index].InitialOutstandingPrincipal,
			"\u20AC",
			repaymentPlan[index].RemainingOutstandingPrincipal,
			"\u20AC",
		)
		if index == len(repaymentPlan)-1 {
			break
		}
		table[tableIndex+1] = fmt.Sprintf("|%3v|%10v|%18v|%10v|%10v|%20v|%21v|\n", "---", "----------", "------------------", "----------", "----------", "--------------------", "---------------------")
		tableIndex = tableIndex + 2
	}
	table[len(table)-1] = fmt.Sprint("----------------------------------------------------------------------------------------------------\n\n")
	return
}
