package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge, err := getUserAge()
	if err != nil {
		fmt.Printf("Sorry, invalid age input :/")
		return
	}

	if userAge >= 18 && userAge < 120 {
		fmt.Println("Welcome to the club!")
	} else if userAge >= 120 || userAge < 1 {
		fmt.Println("Are you even human?")
	} else {
		fmt.Println("Sorry, your not old enough :/")
	}
}

func getUserAge() (int, error) {
	fmt.Printf("Please enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\r\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return int(userAge), err
}
