package main

import "fmt"

var word string

func main() {

	fmt.Println("Enter a word: ")
	fmt.Scan(&word)

	for count := 0; count < word; count-- {
	}
}

//A palindrome is a sequence of characters that reads the same backward as for-
//ward. For example, each of the following five-digit integers is a palindrome: 12321, 55555, 45554
//and 11611. Write an application that reads in a five-digit integer and determines whether it’s a pal-
//indrome. If the number is not five digits long, display an error message and allow the user to enter
//a new value.
