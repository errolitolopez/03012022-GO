package main

import "fmt"

// import "fmt"

// func main() {

// 	var productNames = [4]string{"Book"}

// 	productNames[2] = "Carpet"

// 	fmt.Println(productNames)
// 	fmt.Println(productNames[0])

// 	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
// 	fmt.Println(prices)

// 	fmt.Printf("Name: %v, Price: %.2f\n", productNames[0], prices[3])

// 	// featuredPrices := prices[1:3]
// 	featuredPrices := prices[1:] // starts at index 1 to index 3
// 	fmt.Println(featuredPrices)
// 	featuredPrices[0] = 59.99

// 	highlightedPrices := featuredPrices[:1] // starts at index 0 to index 1
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3] // selects everyting

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Printf("prices: %v\n", prices)

// 	// fmt.Printf("prices: %v\n", prices[0:1])

// 	updatedPrices := append(prices, 5.99)
// 	fmt.Printf("updated prices: %v\n", updatedPrices)

// 	updatedPrices = append(updatedPrices, 6.99)
// 	fmt.Printf("updated prices: %v\n", updatedPrices)

// 	fmt.Printf("prices: %v\n", prices)
// }

// Unpack lists
func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[:1])

	prices[1] = 9.99
	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)
}
