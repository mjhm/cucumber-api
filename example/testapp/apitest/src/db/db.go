// Package that interfaces with the sqlite database
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"user"
)

// openDB opens the given sqlite database by filepath.
func OpenDB(filename string) (*sql.DB, error) {
	return sql.Open("sqlite3", filename)
}

// CommitUser commits the given user to the given database
func CommitUser(db *sql.DB, user user.User) error {
	stmt, err := db.Prepare("insert into users values(NULL,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

// CreateSkelDB creates a skeleton database with one user table and no users.
func CreateSkelDB() error {
	db, err := OpenDB("start.db")
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
