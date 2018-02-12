package cli

import (
	"fmt"
	"log"
	"os"

	plangenerator "github.com/thorion3006/plan-generator"
)

// Creates a file called "Repayment Details.txt" on first call in the root directory and appends the file on subsequent calls.
// It takes two parameters, a slice of strings containing repayment plan and a reference to the LoanDetails struct corresponding to the repayment plan.
func appendFile(table []string, loan *plangenerator.LoanDetails) {
	f, err := os.OpenFile("Repayment details.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	loanDetails := fmt.Sprintf("%s:\n%s: %.2f%s\n%s: %.2f%s\n%s: %d%s\n%s: %.2f%s\n%s: %v\n%s\n",
		"Loan Details",
		"Loan Amount",
		loan.LoanAmount,
		"\u20AC",
		"Nominal Interest Rate",
		loan.NominalRate,
		"%",
		"Duration",
		loan.Duration,
		" months",
		"Annuity",
		loan.Annuity,
		"\u20AC",
		"Start-Date",
		loan.StartDate.Format("02-01-2006"),
		"Repayment Plan:",
	)

	if _, err := f.WriteString(loanDetails); err != nil {
		log.Fatal(err)
	}

	for _, row := range table {
		if _, err := f.WriteString(row); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("File update success!")
}
