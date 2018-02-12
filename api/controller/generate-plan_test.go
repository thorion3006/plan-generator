package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/andreyvit/diff"
	"github.com/thorion3006/plan-generator"
)

var capturedData []byte

func TestHandleRepaymentPlan(t *testing.T) {
	// arrange
	bodyReader := strings.NewReader("{\"loanAmount\":	5000,\"nominalRate\":	5,\"duration\":	2,\"startDate\":	\"2018-01-01T00:00:00Z\"}")
	req := httptest.NewRequest("POST", "/generate-plan", bodyReader)
	var w mockResponseWriter
	w.header = make(map[string][]string)
	var controller generatePlanController
	date, _ := time.Parse("02-01-2006", "01-01-2018")
	loan := plangenerator.LoanDetails{LoanAmount: 5000, Duration: 2, NominalRate: 5, StartDate: date}
	repaymentPlan, _ := plangenerator.RepaymentPlan(&loan)
	repaymentPlanJSON, _ := json.Marshal(repaymentPlan)

	// act
	controller.handleRepaymentPlan(w, req)

	// assert
	if w.header.Get("Content-Type") != "application/json" {
		t.Error("Missing or incorrect Content-Type header")
	}

	if a, e := string(capturedData), (string(repaymentPlanJSON) + "\n"); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
	}
}

// To mimic the functions of http.ResponseWriter
type mockResponseWriter struct {
	header http.Header
}

func (w mockResponseWriter) Header() http.Header {
	return w.header
}

func (w mockResponseWriter) Write(data []byte) (int, error) {
	capturedData = data[:]
	return len(data), nil
}

func (w mockResponseWriter) WriteHeader(code int) {

}
