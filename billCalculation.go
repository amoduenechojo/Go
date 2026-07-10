package main

import (
	"fmt"
	"time"
)

func billCalculation(customer string, cashier string, discountPercent float64, cart []CartItem) {
	var amountGivenToCashierByCustomer float64
	var subTotal float64

	now := time.Now()

	fmt.Println("SEMICOLON BRANCH")
	fmt.Println("MAIN BRANCH")
	fmt.Println("LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.")
	fmt.Println("03293826343")

	fmt.Println("Date:", now.Format("02-Jan-06 3:04:05 pm"))
	fmt.Println("Cashier:", cashier)
	fmt.Println("Customer Name:", customer)

	fmt.Println("=========================================================================================")
	rowFormat := "%-20s %-10s %-15s %-15s\n"
	fmt.Printf(rowFormat, "ITEM", "QTY", "PRICE", "TOTAL (NGN)")
	fmt.Println("-----------------------------------------------------------------------------------------")

	for index := 0; index < len(cart); index++ {
		item := cart[index]

		itemTotal := float64(item.Quantity) * item.Price
		subTotal += itemTotal
		fmt.Printf("%-20s %-10d %-15.2f %-15.2f\n", item.Name, item.Quantity, item.Price, itemTotal)
	}
	fmt.Println("-----------------------------------------------------------------------------------------")

	discountAmount := (discountPercent / 100.0) * subTotal
	VATAmount := 0.175 * subTotal
	billTotal := (subTotal - discountAmount) + VATAmount

	fmt.Printf("%-48s %.2f\n", "Sub Total:", subTotal)
	fmt.Printf("%-48s %.2f\n", "Discount:", discountAmount)
	fmt.Printf("%-48s %.2f\n", "VAT @ 17.50%:", VATAmount)
	fmt.Println("=========================================================================================")
	fmt.Printf("%-48s %.2f\n", "Bill Total:", billTotal)
}
