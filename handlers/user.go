package handlers

import (
	"net/http"
	"strconv"
	"todo_app/data"
	"todo_app/models"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := data.GetUsers()
	return c.JSON(http.StatusOK, users)
}

func AddUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	user.ID = len(data.GetUsers()) + 1
	data.AddUser(user)
	return c.JSON(http.StatusCreated, user)
}

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := data.GetUserByID(id)
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}