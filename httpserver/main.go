package main

import (
	"fmt"
	"log"
	"net/http"
)

func helthHandler() {

}

func formHandler() {

}

func main() {
	fileserver := http.FileServer((http.Dir(("./static"))))
	http.Handle("/", fileserver)

	http.HandleFunc("/health", helthHandler)
	htttp.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
