package controllers

import (
	"golang-todo-v2-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllTodos(c echo.Context) error {
	todos := []models.Todo{}
	models.GetAllTodos(&todos)
	return c.JSON(http.StatusOK, todos)
}

func createTodo(c echo.Context) error {
	t := new(models.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	models.CreateTodo(t)
	return c.JSON(http.StatusOK, t)
}

func deleteTodo(c echo.Context) error {
	id := c.Param("id")
	t := new(models.Todo)
	models.DeleteTodo(t, id)
	return c.NoContent(http.StatusNoContent)
}
