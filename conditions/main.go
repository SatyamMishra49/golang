package main

import "fmt"

func cdts() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	var y = 20

	if y < 0 {
		fmt.Println("y is negative")
	}
	if y >= 18 && y < 65 {
		fmt.Println("y is an adult")
	}
}

func altfunc() {
	var x int = 15

	if x < 10 {
		fmt.Println("x is less than 10")
	} else if x >= 10 && x < 20 {
		fmt.Println("x is between 10 and 20")
	} else {
		fmt.Println("x is 20 or more")
	}
}

func main() {
	cdts()
	altfunc()
	fmt.Println("Conditions executed successfully")
}
