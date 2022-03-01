package main

import "fmt"

func main() {

	const PI = 3.14
	radius := 5

	var circumference = 2 * PI * float64(radius)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f.", radius, circumference)

}
