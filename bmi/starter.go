package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("-----------------------------")
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 32)
	height, _ := strconv.ParseFloat(heightInput, 32)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)
}
