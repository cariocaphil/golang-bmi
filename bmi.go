package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI calculator")
	fmt.Println("--------------")
	fmt.Println("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')
	fmt.Println("Please enter your height (m): ")
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

	fmt.Println(weight)
	fmt.Println(height)

}
