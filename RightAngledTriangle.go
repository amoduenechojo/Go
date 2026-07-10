package main

import "fmt"

var triangleLength int

func triangle() {

	fmt.Println("Enter a number of length of your triangle: ")
	fmt.Scan(&triangleLength)

	for count := 1; count <= triangleLength; count++ {
		for index := 1; index <= count; index++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func main() {
	triangle()
}
