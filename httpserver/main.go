package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler() {

}

func formHandler() {

}

func main() {
	fileserver := http.FileServer((http.Dir(("./static"))))
	http.Handle("/", fileserver)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
