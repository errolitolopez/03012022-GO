package main

import "fmt"

func main() {

	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Printf("total: %v\n", total)

	nums := []int{10, 10}
	total = sum(nums...)
	fmt.Printf("total: %v\n", total)
}

func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
