package main

import "fmt"

// const pi float64 = 3.141522539

// Format Printing
// func main() {
// 	const x float64 = 100.926748729
// 	y := 100
// 	Name := "Juan Dela Cruz"
// 	var isBool bool = false
// 	fmt.Printf("%.2f \n", x)
// 	fmt.Printf("%T \n", y)
// 	fmt.Printf("%T \n", Name)
// 	fmt.Printf("%T \n", isBool)
// 	fmt.Printf("%t \n", isBool)
// 	fmt.Printf("%d \n", y)
// }

// Values and Variables
// func main() {
// 	fmt.Println(pi)
// 	var Name string = "Juan Dela Cruz"
// 	fmt.Println(Name)
// 	fmt.Println(len(Name))
// 	var num int = 1
// 	fmt.Println(num)
// 	var (
// 		varA = 1
// 		varB = 2
// 	)
// 	fmt.Println(varA, varB)

// 	var firstName string = "Juan"
// 	lastName := "Dela Cruz"

// 	fmt.Println(firstName + " " + lastName)

// 	currentYear := time.Now().Year()
// 	birthYear := time.Date(2000, time.December, 12, 0, 0, 0, 0, time.UTC).Year()
// 	age := currentYear - birthYear
// 	nextYear := currentYear + 1

// 	fmt.Println(currentYear)
// 	fmt.Println(birthYear)
// 	fmt.Println(age)
// 	fmt.Println(nextYear)
// }

// Operators
// func main() {
// 	x, y := 10, 4
// 	isTrue := true
// 	isFalse := false
// 	// var x float32 = 10
// 	// var y float32 = 4
// 	fmt.Println("x + y = ", x+y)
// 	fmt.Println("x - y = ", x-y)
// 	fmt.Println("x * y = ", x*y)
// 	fmt.Println("x / y = ", x/y)
// 	fmt.Println("x % y = ", x%y)
// 	fmt.Println("x >= y = ", x >= y)
// 	fmt.Println("x <= y = ", x <= y)
// 	fmt.Println("x == y = ", x == y)
// 	fmt.Println("x != y = ", x != y)
// 	fmt.Println(isTrue)
// 	fmt.Println(isFalse)
// 	fmt.Println(isTrue && isFalse)
// 	fmt.Println(isTrue || isFalse)
// 	fmt.Println(!isFalse)
// }

// Ifs
// func main() {

// 	age := 18
// 	if age < 18 {
// 		fmt.Println("You can't vote!")
// 	} else {
// 		fmt.Println("You can vote!")
// 	}

// 	switch age {
// 	case 16:
// 		fmt.Println("Prepare for college!")
// 	case 18:
// 		fmt.Println("Get a Job!")
// 	default:
// 		fmt.Println("Are you even alive?")
// 	}
// }

// Loops
func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
