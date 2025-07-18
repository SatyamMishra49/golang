package main

import "fmt"

func PublicFunc() {
	fmt.Println("This is a public function")
}

func privateFunc() {
	fmt.Println("This is a private function")
}

func main() {
	age := 12
	name := "Alice"

	fmt.Println("Age:", age, "Name:", name)

	fmt.Printf("age is %d\n", age)
	fmt.Printf("name is %s\n", name)

	fmt.Printf("Type of name is : %T\n", name) // %T is used to get the type of a variable
}
