package controllers

import (
	"golang-todo-v2-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createTodo(c echo.Context) error {
	t := new(models.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	models.CreateTodo(t)
	return c.JSON(http.StatusOK, t)
}
