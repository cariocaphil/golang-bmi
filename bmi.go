package main

func main() {

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	printBMI(bmi)

}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
