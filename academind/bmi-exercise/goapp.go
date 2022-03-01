package main

import "exercise/bmi/info"

func main() {
	info.PrintWelcome()
	weight, height := getUserMetrics()
	bmi := calculateBMI(weight, height)
	outputBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
