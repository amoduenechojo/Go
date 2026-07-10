package main

import "fmt"

type CartItem struct {
	Name     string
	Quantity int
	Price    float64
}

func main() {
	var customersName string
	var cashierName string
	var customersDiscount float64

	var cart []CartItem

	fmt.Println("Enter customer name: ")
	fmt.Scan(&customersName)

	for {

		var itemBought string
		var itemPieces int
		var unitPrice float64
		var moreItem string

		fmt.Println("Enter item bought: ")
		fmt.Scan(&itemBought)

		fmt.Println("How many pieces did the customer buy?: ")
		fmt.Scan(&itemPieces)

		fmt.Print("How much per unit? ")
		fmt.Scan(&unitPrice)

		cart = append(cart, CartItem{
			Name:     itemBought,
			Quantity: itemPieces,
			Price:    unitPrice,
		})

		fmt.Print("Add more Items? (yes/no): ")
		fmt.Scan(&moreItem)

		if moreItem != "yes" {
			break
		}
	}

	fmt.Println("Cashier's name: ")
	fmt.Scan(&cashierName)

	fmt.Println("How much discount will the customer get?")
	fmt.Scan(&customersDiscount)

	billCalculation(customersName, cashierName, customersDiscount, cart)

	fmt.Println("=========================================================================================")

	fmt.Println("How much did the customer give to you?")
	fmt.Scan(&amountGivenToCashierByCustomer)

	receiptGeneration(customersName, cashierName, discountPercent, cart, amountGivenToCashierByCustomer)
}
