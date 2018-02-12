package cli

import (
	"testing"
	"time"

	plangenerator "github.com/thorion3006/plan-generator"
)

func TestPrintTable(t *testing.T) {
	// arrange
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	loan := plangenerator.LoanDetails{LoanAmount: 1000, NominalRate: 12, Duration: 2, StartDate: date}
	repaymentPlan, err := plangenerator.RepaymentPlan(&loan)
	if err != nil {
		return
	}
	table := make([]string, 8)
	table[0] = "----------------------------------------------------------------------------------------------------\n"
	table[1] = "|  #|      Date| Annuity (Borrower| Principal|  Interest| Initial Outstanding|Remaining Outstanding|\n"
	table[2] = "|   |          |   Payment Amount)|          |          |           Principal|            Principal|\n"
	table[3] = "|---|----------|------------------|----------|----------|--------------------|---------------------|\n"
	table[4] = "|  1|01-01-2018|           507.51€|   497.51€|    10.00€|            1000.00€|              502.49€|\n"
	table[5] = "|---|----------|------------------|----------|----------|--------------------|---------------------|\n"
	table[6] = "|  2|01-02-2018|           507.51€|   502.49€|     5.02€|             502.49€|                0.00€|\n"
	table[7] = "----------------------------------------------------------------------------------------------------\n\n"

	// act
	result := printTable(repaymentPlan)

	// assert
	for index, value := range result {
		if value != table[index] {
			t.Errorf("Expected: %sGot: %s", table[index], value)
		}
	}
}
