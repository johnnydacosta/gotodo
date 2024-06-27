package main

import (
	"fmt"
	"net/http"

	"github.com/johnnydacosta/gotodo/todo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", GetTodos)
	e.POST("/", AddTodo)
	e.GET("/:id", GetTodo)
	e.PUT("/:id", UpdateTodo)
	e.DELETE("/:id", DeleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, todo.FindAll())
}

func GetTodo(c echo.Context) error {
	var form todo.Todo

	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
	}

	return c.JSON(http.StatusOK, todo.FindById(form.ID))
}

func AddTodo(c echo.Context) error {
	var form todo.CreateTodoDTO

	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
	}

	return c.JSON(http.StatusCreated, todo.Create(form))
}

func UpdateTodo(c echo.Context) error {
	var form todo.UpdateTodoDTO

	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
	}

	t := todo.Update(form)

	return c.JSON(http.StatusOK, t)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	todo.Delete(id)
	return c.JSON(http.StatusNoContent, nil)
}
