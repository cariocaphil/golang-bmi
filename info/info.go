package info

import (
	"fmt"
)

const mainTitle = "BMI calculator"
const separator = "--------------"
const WeightPrompt = "Please enter your weight (kg): "
const HeightPrompt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
