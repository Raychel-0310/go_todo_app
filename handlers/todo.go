package handlers

import (
	"net/http"
	"strconv"
	"todo_app/data"
	"todo_app/models"

	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	tasks := data.GetTasks()
	return c.JSON(http.StatusOK, tasks)
}

func AddTask(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return err
		}

	task.ID = len(data.GetTasks()) + 1
	data.AddTask(task)
	return c.JSON(http.StatusCreated, task)
}

func UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var task models.Task
	if err := c.Bind(&task); err != nil {
		return err
	}

	data.UpdateTask(id, task.Task, task.Done)
	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	data.DeleteTask(id)
	return c.NoContent(http.StatusNoContent)
}