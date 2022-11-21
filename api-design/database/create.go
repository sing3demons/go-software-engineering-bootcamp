package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	url := "host=127.0.0.1 port=5432 user=sing password=12345678 dbname=goapi sslmode=disable"
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	log.Println("ok")
}
