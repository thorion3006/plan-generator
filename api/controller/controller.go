package controller

import "net/http"

var generatePlan generatePlanController

// Register the controllers to the server.
func Setup() {
	generatePlan.registerRoute()
	http.Handle("/", http.NotFoundHandler())
}
