package main

import "fmt"

func main() {
	age := 26
	fmt.Println(age)

	// var myAge *int>
	// myAge = &age

	myAge := &age

	fmt.Println(myAge)
	fmt.Println(*myAge) // dereferencing

	*myAge = 25

	println(*myAge)
	println(age)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)

	println(age)

	println(myAge)
}

func double(number *int) int {
	result := *number * 2
	*number = 100
	return result
}
