package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/cariocaphil/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {
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

	return weight, height
}
