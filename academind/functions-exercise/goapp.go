package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	a := 5
	b := 10
	sum := add(a, b)
	printNumber(sum)
	fmt.Println("I'm executing inside a function!")
	c, d := generateRandomNumbers()
	printNumber(c)
	printNumber(d)
	printNumber(add(c, d))
}

func add(a int, b int) (sum int) {
	sum = a + b
	return
}

func printNumber(number int) {
	fmt.Printf("The number is %v\n", number)
}

func generateRandomNumbers() (r1 int, r2 int) {
	randomNumber1, _ := rand.Int(rand.Reader, big.NewInt(10))
	randomNumber2, _ := rand.Int(rand.Reader, big.NewInt(100))
	return int(randomNumber1.Int64()), int(randomNumber2.Int64())
}
