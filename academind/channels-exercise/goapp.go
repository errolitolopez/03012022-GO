package main

import "fmt"

func main() {

	// messages := make(chan string)
	// go func() {
	// 	messages <- "ping"
	// }()
	// msg := <-messages

	// fmt.Printf("msg: %v\n", msg)

	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	v := <-ch

	fmt.Printf("v: %v\n", v)
}
