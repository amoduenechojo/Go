package ChapterFour

import "fmt"

func calculatingTotalGasUsed() {

	var milesDriven int = 0
	var gallonsUsed int = 0
	var totalMilesDriven int = 0
	var totalGallonsUsed int = 0

	fmt.Println("Enter the amount of miles driven: ")
	fmt.Scan(&milesDriven)

	if milesDriven < 1 {
		fmt.Println("Enter the amount of miles driven: ")
	}

	fmt.Println("Enter the amount of gallons used: ")
	fmt.Scan(&gallonsUsed)

	if gallonsUsed == 0 {
		fmt.Println("Enter the amount of gallons used: ", "Gallon used cannot be zero.")
	}

	totalMilesDriven += milesDriven
	totalGallonsUsed += gallonsUsed

	milesPerGallon := float64(totalMilesDriven) / float64(totalGallonsUsed)

}

func main() {
	calculatingTotalGasUsed()
}
