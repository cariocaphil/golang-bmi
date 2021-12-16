package main

import (
	"fmt"
)

func main() {

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	fmt.Printf("Your BMI is: %.2f ", bmi)

}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
