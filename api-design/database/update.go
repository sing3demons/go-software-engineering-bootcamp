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

	stmt, err := db.Prepare("UPDATE users SET name=$2, age=$3 WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare query all users statement", err)
	}

	if _, err := stmt.Exec(3, "kp", 22); err != nil {
		log.Fatal("error execute update", err)
	}

	fmt.Println("update success")
}
