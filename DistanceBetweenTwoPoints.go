package main

import "fmt"

var x1 int
var x2 int

var y1 int
var y2 int

func main() {
	fmt.Print("Enter a number: ")
	fmt.Scan(&x1)

	fmt.Print("Enter a number: ")
	fmt.Scan(&x2)

	fmt.Print("Enter a number: ")
	fmt.Scan(&y1)

	fmt.Print("Enter a number: ")
	fmt.Scan(&y2)

	if x1 == x2 {
		fmt.Println("The line is perpendicular to the x-axis(vertical line")
	}

	if y1 == y2 {
		fmt.Println("The line is perpendicular to the y-axis(vertical line")
	} else {
		fmt.Println("The line is not perpendicular to any axis.")
	}
}
