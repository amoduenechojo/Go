package main

import "fmt"

var earning float32
var taxRate float32
var totalTax float32
var citizensName string

func taxCalculation() {

	for count := 1; count <= 3; count++ {

		fmt.Println("Enter citizen's name: ")
		fmt.Scan(&citizensName)

		fmt.Println("Enter how much you earn: ")
		fmt.Scan(&earning)

		if earning <= 30000 {
			taxRate = 0.15 * earning
		}

		if earning > 30000 {
			taxRate = (0.15 * earning) + (0.20*earning - 30000)
		}

		totalTax = taxRate

		fmt.Printf("Citizen: %s | Total Tax: $%.2f\n\n", citizensName, totalTax)
	}

}

func main() {
	taxCalculation()
}
