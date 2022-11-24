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

	e.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if username == "admin" && password == "1234" {
			return true, nil
		}

		return false, nil
	}))

	e.GET("/users", getUserHandler)
	e.POST("/users", createUserHandler)
	e.GET("/users/:id", getUserHandlerByID)
	e.DELETE("/users/:id", deleteUserHandlerByID)
	e.PUT("/users/:id", updateUserHandler)

	e.Logger.Fatal(e.Start(":2565"))
}

func updateUserHandler(c echo.Context) error {
	id := c.Param("id")
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	stmt, err := db.Prepare("UPDATE users SET name=$2, age=$3 WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}
	_, err = stmt.Exec(id, u.Name, u.Age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}
	return c.JSON(http.StatusNoContent, "Update")
}

func deleteUserHandlerByID(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("DELETE FROM users where id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't scan user" + err.Error()})
	}
	return c.JSON(http.StatusNoContent, "Delete")
}

func getUserHandlerByID(c echo.Context) error {
	var user User
	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}
	
	id := c.Param("id")
	row := stmt.QueryRow(id)

	err = row.Scan(&user.ID, &user.Name, &user.Age)
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, user)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}

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
