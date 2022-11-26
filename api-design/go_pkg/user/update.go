package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateUserHandler(c echo.Context) error {
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
