package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sing3demons/echoapisql/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	user.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", healthHandler)

	e.Use(middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
		if username == "admin" && password == "1234" {
			return true, nil
		}

		return false, nil
	}))

	e.GET("/users", user.GetUserHandler)
	e.POST("/users", user.CreateUserHandler)
	e.GET("/users/:id", user.GetUserHandlerByID)
	e.DELETE("/users/:id", user.DeleteUserHandlerByID)
	e.PUT("/users/:id", user.UpdateUserHandler)

	go func() {
		if err := e.Start(":2565"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	<-shutdown
	fmt.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("server stop")
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
