package main

import (
	"fmt"

	"github.com/thorion3006/plan-generator/api"
	"github.com/thorion3006/plan-generator/cli"
)

func main() {
	fmt.Println("Welcome to Loan repayment Calculator!")
	fmt.Println()
	fmt.Println("Menu")
	fmt.Println("1. Start the Command Line Interface Application.")
	fmt.Println("2. Start the web Api")
	fmt.Print("Please select an option:(1/2) ")
	userResponse := cli.UserResponseProcessor()
	switch userResponse {
	case "1":
		cli.ClearScreen()
		cli.Start()
	case "2":
		cli.ClearScreen()
		api.Start()
	default:
		fmt.Println("Please try again with a valid input.")
	}
}
