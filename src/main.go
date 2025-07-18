package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var name string = "John Doe"
	fmt.Println("Name:", name)

	var money int = 100
	var currency = 12000
	fmt.Println("Money:", money)
	fmt.Println("Currency:", currency)

	var dimension float64 = 3.14
	fmt.Println("Dimension:", dimension)

	var isTrue bool = true
	isTrue = false // Reassigning the value
	fmt.Println("Is True:", isTrue)

	const pi = 3.17
	fmt.Println("Pi:", pi)

	person := "Man"
	fmt.Println("Person:", person)

	// Public and private variables
	var PublicVar = "This is public"
	var privateVar = "This is private"
	fmt.Println("Public Variable:", PublicVar)
	fmt.Println("Private Variable:", privateVar)
}
