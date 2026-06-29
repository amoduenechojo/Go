package main

import "fmt"

var balanceAtTheBeginningOfTheMonth int
var totalItemCharged int
var customersAppliedCredit int
var customersCreditLimit int

func usersCardInformation() {

	var accountNumber string

	fmt.Println("Enter your account number: ")
	fmt.Scan(&accountNumber)

	fmt.Println("What was your balance at the beginning of the month: ")
	fmt.Scan(&balanceAtTheBeginningOfTheMonth)

	fmt.Println("Enter total item charged by the customer this month: ")
	fmt.Scan(&totalItemCharged)

	fmt.Println("Enter the total of all credit applied to the customers account this month: ")
	fmt.Scan(&customersAppliedCredit)

	fmt.Println("Enter customers credit limit: ")
	fmt.Scan(&customersCreditLimit)
}

func balance() {

	newCustomersBalance := balanceAtTheBeginningOfTheMonth + totalItemCharged - customersAppliedCredit
	fmt.Println("Your new balance is: ", newCustomersBalance)

	if newCustomersBalance > customersCreditLimit {
		fmt.Println("Credit limit exceeded")
	}
}

func main() {
	usersCardInformation()
	balance()
}
