package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Err struct {
	Message string `json:"message"`
}

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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", healthHandler)

	e.GET("/users", getUserHandler)
	e.POST("/users", createUserHandler)

	e.Logger.Fatal(e.Start(":2565"))
}

func getUserHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("can't prepare query all users statement", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func createUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", u.Name, u.Age)

	err = row.Scan(&u.ID)
	if err != nil {
		log.Fatal("can't insert data", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"id": u.ID,
	})
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
