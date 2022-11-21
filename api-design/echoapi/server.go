package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
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

var users []User = []User{{ID: 1, Name: "sing", Age: 21}}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)
	users := e.Group("api")
	users.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if username == "admin" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))
	{
		users.GET("/users", getUserHandler)
		users.POST("/users", createUserHandler)
	}

	e.Logger.Fatal(e.Start(":2565"))

}

func getUserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func createUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	id := len(users) + 1
	u.ID = id

	users = append(users, u)
	return c.JSON(http.StatusCreated, users)
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
