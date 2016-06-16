// Package that interfaces with the sqlite database
package db

import (
  _ "github.com/mattn/go-sqlite3"
  "database/sql"

)

func openDB(filename string) (*sql.DB, error) {
  return sql.Open("sqlite3", filename)
}

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
    age integer,
    friends text not null,
  )
  `
}
