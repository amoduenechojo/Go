package main

import "fmt"

var firstNumber int
var secondNumber int

func main() {
	fmt.Println("Enter the first number:")
	fmt.Scan(&firstNumber)

	fmt.Println("Enter the second number:")
	fmt.Scan(&secondNumber)

	if firstNumber == secondNumber {
		fmt.Println(0)
	}

	if firstNumber > secondNumber {
		fmt.Println(1)
	} else if secondNumber > firstNumber {
		fmt.Println(-2)
	}
}
