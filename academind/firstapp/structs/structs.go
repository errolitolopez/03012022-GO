package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdDate time.Time
}

func (user *User) outPutDetails() {
	fmt.Printf("My name is: %v %v (born on %v)\n", user.firstName, user.lastName, user.birthdate)
}

func main() {
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")

	newUser = NewUser(firstName, lastName, birthdate)
	newUser.outPutDetails()
}

var reader = bufio.NewReader(os.Stdin)

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\r\n", "", -1)
	return cleanedInput
}

func NewUser(fName string, lName string, bDate string) *User {
	return &User{
		fName,
		lName,
		bDate,
		time.Now(),
	}
}
