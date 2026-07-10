package main

import "fmt"

var digit int

func main() {

	fmt.Print("Enter a number: ")
	fmt.Scan(&digit)

	for count := 0; count < digit; count-- {
		if digit < 0 {
			fmt.Printf("Number cannot be less than zero.")
		}

		count *= digit
		fmt.Println(count)
	}
}
