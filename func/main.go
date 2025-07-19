package main

import "fmt"

/*
func name(parameter-list) (result-list) {
}
*/

func myfunc(fname string, lname string) string {
	name := fname + " " + lname
	return name
}

func mfunc(fname string, lname string) (string, error) {
	name := fname + " " + lname
	return name, nil
}

//Anonymous functions
func addOne() func() int {
	var x int

	return func() int {
		x++
		return x + 1
	}
}

func main() {

	result := myfunc("John", "Doe")
	fmt.Println(result)

	fullname, err := mfunc("Jane", "Smith")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(fullname)
	}

	increment := addOne()
	fmt.Println(increment())
}
