package main

// Test server app from cucumber-api

import (
	"database/sql"
	"db"
	"fmt"
	"log"
	"net/http"
	"os"
)

var dbase *sql.DB

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

// Giving main a flag of '--create' will simply create a skeleton database and exit
func main() {
	if len(os.Args) > 1 && os.Args[1] == "--create" {
		err := db.CreateSkelDB()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	dbase, err := db.OpenDB("start.db")
	setGlobalDatabase(dbase)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/users", handleUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setGlobalDatabase(d *sql.DB) {
	dbase = d
}
