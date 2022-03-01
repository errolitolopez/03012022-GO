package main

import "fmt"

func main() {
	x := 10
	y := 5
	fmt.Println(&x)
	fmt.Println(&y)
	changeValue(&x)
	fmt.Println(x)

}

func changeValue(x *int) {
	*x = 7
}
