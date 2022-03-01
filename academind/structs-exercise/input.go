package main

import (
	"bufio"
	"exercise/structs/info"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getProduct() *Product {
	idInput := readUserInput(reader, info.IdPrompt)
	idValue, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Printf("error: \"id\" message: %v", err.Error())
		os.Exit(0)
	}

	nameInput := readUserInput(reader, info.NamePrompt)
	descriptionInput := readUserInput(reader, info.DescriptionPrompt)

	priceInput := readUserInput(reader, info.PricePrompt)
	priceValue, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		fmt.Printf("error: \"price\" message: %v", err.Error())
		os.Exit(0)
	}

	return NewProduct(idValue, nameInput, descriptionInput, priceValue)
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	return userInput
}
