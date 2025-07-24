package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("./dist")
	check(err)
	fmt.Println("Data read from file:", string(data))

	fmt.Println("Main Func Running...")
}
