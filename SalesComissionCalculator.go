package main

import "fmt"

var salesItemForThePreviousWeek float32

func salesPersonIncome() {
	var totalGrossSales float32

	for {
		fmt.Println("Enter your sales item for the previous week: ")
		fmt.Scan(&salesItemForThePreviousWeek)

		if salesItemForThePreviousWeek == 0 {
			break
		}

		totalGrossSales += salesItemForThePreviousWeek
	}

	salesBonusPay := 200 + (0.09 * salesItemForThePreviousWeek)

	fmt.Printf("Total Gross Sales: $%.2f\n", totalGrossSales)
	fmt.Printf("Total Earnings (Base $200 + 9%): $%.2f\n", salesBonusPay)
}

func main() {
	salesPersonIncome()
}
