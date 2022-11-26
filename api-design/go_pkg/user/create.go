package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
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
