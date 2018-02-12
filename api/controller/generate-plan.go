package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thorion3006/plan-generator"
)

type generatePlanController struct{}

// Add the /generate-plan route to the controller list.
func (g *generatePlanController) registerRoute() {
	http.HandleFunc("/generate-plan", g.handleRepaymentPlan)
}

// Handling the post action on the /generate-plan route
func (g *generatePlanController) handleRepaymentPlan(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loan plangenerator.LoanDetails
	err := dec.Decode(&loan)
	if err != nil {
		response := "Error: Invalid data type provided. Valid data types are float for Loan Amount and Rate, int for Duration and date object for Start-Date."
		http.Error(w, response, 400)
		return
	}
	repaymentPlan, err := plangenerator.RepaymentPlan(&loan)
	if err != nil {
		response := fmt.Sprint(err)
		http.Error(w, response, 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(repaymentPlan)
}
