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

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "sing", 25)

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}

	fmt.Println("insert todo success id")
}
