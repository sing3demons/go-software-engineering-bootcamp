package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUserHandlerByID(c echo.Context) error {
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
