package main

import "fmt"

func main() {

	var currentNumber int

	largestNumber := 0
	secondLargestNumber := -1

	for count := 1; count <= 10; count++ {
		fmt.Print("Enter number: ")
		fmt.Scan(&currentNumber)

		if currentNumber > largestNumber {
			secondLargestNumber = largestNumber
			largestNumber = currentNumber

		} else if currentNumber > secondLargestNumber {
			secondLargestNumber = currentNumber
		}

	}

	fmt.Println("The largest number is", largestNumber)
	fmt.Println("The second largest number is", secondLargestNumber)
}
