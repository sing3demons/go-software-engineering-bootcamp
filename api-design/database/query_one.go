package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "host=127.0.0.1 port=5432 user=sing password=12345678 dbname=goapi sslmode=disable"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one user statement", err)
	}

	rowId := 1

	row := stmt.QueryRow(rowId)

	if err := row.Err(); err != nil {
		log.Fatal("can't query all users", err)
	}

	var id, age int
	var name string
	if err := row.Scan(&id, &name, &age); err != nil {
		log.Fatal("can't scan row into variable", err)
	}
	fmt.Println(id, name, age)
}
