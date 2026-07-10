package main

import "fmt"


var billTotal float64
var amountGivenToCashierByCustomer float64

func receiptGeneration(customer string, cashier string, discountPercent float64, cart []CartItem, amountGivenToCashierByCustomer float64){

	if amountGivenToCashierByCustomer < billTotal {

		fmt.Println("======================================================")
		fmt.Printf("%-48s Insufficient Balance", )
		fmt.Println("=========================================================================================")
	}

	else{
		billCalculation(customer string, cashier string, discountPercent float64, cart []CartItem)
	}

}
