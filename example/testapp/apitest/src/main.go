package main

// Test server app from cucumber-api

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
