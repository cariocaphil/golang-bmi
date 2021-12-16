package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/cariocaphil/bmi/info"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Println(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Println(info.WeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// Save that user input in variables
	if runtime.GOOS == "windows" {
		weightInput = strings.Replace(weightInput, "\r\n", "", -1)
		heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	} else {
		weightInput = strings.Replace(weightInput, "\n", "", -1)
		heightInput = strings.Replace(heightInput, "\n", "", -1)
	}

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	fmt.Println(reflect.TypeOf(weight))

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is: %.2f ", bmi)

}
