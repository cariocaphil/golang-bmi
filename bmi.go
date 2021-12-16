package main

import (
	"fmt"
)

func main() {

	weight, height := getUserMetrics()

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is: %.2f ", bmi)

}
