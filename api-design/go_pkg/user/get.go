package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserHandler(c echo.Context) error {
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

func GetUserHandlerByID(c echo.Context) error {
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
