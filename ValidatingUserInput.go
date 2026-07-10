package main

import "fmt"

var userInput int

func validateUserInput() {

	fmt.Println("Please enter a number: ")
	fmt.Scan(&userInput)

	for userInput > 1 {
		if userInput != 1 && userInput != 2 {
			fmt.Println("Invalid user input! Try again!")
		}
	}
}

func main() {
	validateUserInput()
}
