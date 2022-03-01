package main

import "fmt"

func main() {

	nextInt := intSeq()
	fmt.Printf("nextInt: %v\n", nextInt())
	fmt.Printf("nextInt: %v\n", nextInt())
	fmt.Printf("nextInt: %v\n", nextInt())
	fmt.Printf("nextInt: %v\n", nextInt())
	fmt.Printf("nextInt: %v\n", nextInt())
	newInt := intSeq()
	fmt.Printf("newInt: %v\n", newInt())
	fmt.Printf("newInt: %v\n", newInt())
	fmt.Printf("newInt: %v\n", newInt())
	fmt.Printf("nextInt: %v\n", nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
