package main

import "fmt"

func main() {

	fmt.Println("N\tN^2\tN^3\tN^4")

	for count := 1; count <= 5; count++ {

		fmt.Printf("%d\t%d\t%d\t%d\n", count, count*count, count*count*count, count*count*count*count)
	}
}
