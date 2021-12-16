package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/cariocaphil/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (float64, float64) {

	weight := getUserInput(info.WeightPrompt)
	height := getUserInput(info.HeightPrompt)
	return weight, height
}

func getUserInput(promptText string) float64 {
	fmt.Println(promptText)
	userInput, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}
	value, _ := strconv.ParseFloat(userInput, 64)
	return value
}
