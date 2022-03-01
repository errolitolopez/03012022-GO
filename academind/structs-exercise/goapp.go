package main

import (
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	id          int
	name        string
	description string
	price       float64
}

func (product *Product) printData() {
	fmt.Printf("Name: %v Description: %v (price starts at %.2f)\n", product.name, product.description, product.price)
}

func main() {

	firstProduct := Product{
		1,
		"iPhone 13 Pro",
		"Smartphone - Apple",
		63989.99,
	}

	secondProduct := *NewProduct(
		2,
		"Samsung Galaxy 20",
		"Smartphone - Samsung",
		49999.99,
	)

	thirdProduct := getProduct()

	firstProduct.printData()
	firstProduct.store()
	secondProduct.printData()
	thirdProduct.printData()

}

func NewProduct(id int, n string, d string, p float64) *Product {
	return &Product{id, n, d, p}
}

func (product *Product) store() {
	file, _ := os.Create(strconv.Itoa(product.id) + ".txt")
	content := fmt.Sprintf("ID: %v\nName: %v\nDescription: %v\nPrice: %.2f",
		product.id,
		product.name,
		product.description,
		product.price,
	)

	file.WriteString(content)
	file.Close()
}
