// Package that interfaces with the sqlite database
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// openDB opens the given sqlite database by filepath.
func openDB(filename string) (*sql.DB, error) {
	return sql.Open("sqlite3", filename)
}


// CreateSkelDB creates a skeleton database with one user table and no users.
func CreateSkelDB() error {
	db, err := openDB("start.db")
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := `
  create table users (
    id integer primary key,
    name text not null,
    age integer
  )
  `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil
	}
	tx.Commit()
	return nil
}
