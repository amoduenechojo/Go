package main

import "fmt"

	var number int
	func largestNumber(){

		count := 10
		largestNumber := 0

		for count > 10{

			if(count != 10){
				fmt.Println("Enter a number: ")
				fmt.Scan(&number)
			}

			if(number > largestNumber){
				largestNumber = number
			}
			count--
		}

		fmt.Println("The largest number is", largestNumber)

	func main(){
		largestNumber()
	}
}

