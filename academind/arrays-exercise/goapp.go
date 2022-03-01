package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {

	// Practice 1
	hobbies := [3]string{
		"Sports",
		"Cooking",
		"Reading",
	}
	fmt.Println(hobbies)

	// Practice 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// Practice 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// Practice 4
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// Practice 5
	courseGoals := []string{"Learn Go arrays!", "Learn Go slices!"}
	fmt.Println(courseGoals)

	// Practice 6
	courseGoals[1] = "Learn Go structs!"
	fmt.Println(courseGoals)

	courseGoals = append(courseGoals, "Learn Go slices!")
	fmt.Println(courseGoals)

	// Practice 7
	products := []Product{
		{
			1,
			"Book",
			1.99,
		},
		{
			2,
			"Carpet",
			59.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{3, "Keyboard", 69.99}
	products = append(products, newProduct)

	fmt.Println(products)
}
