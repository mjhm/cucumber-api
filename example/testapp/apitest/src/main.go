package main

// Test server app from cucumber-api

import (
	"db"
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--create" {
		err := db.CreateSkelDB()
		if err != nil {
			fmt.Println(err)
		}
	}
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
