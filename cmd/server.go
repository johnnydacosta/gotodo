package main

import (
	"net/http"
	"strconv"

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
	id := c.Param("id")
	return c.JSON(http.StatusOK, todo.FindById(id))
}

func AddTodo(c echo.Context) error {
	form := todo.CreateTodoDTO{
		Name: c.FormValue("name"),
	}

	return c.JSON(http.StatusCreated, todo.Create(form))
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	done, err := strconv.ParseBool(c.FormValue("done"))
	if err != nil {
	}

	form := todo.UpdateTodoDTO{
		Name: c.FormValue("name"),
		Done: done,
	}

	t := todo.Update(id, form)

	return c.JSON(http.StatusOK, t)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	todo.Delete(id)
	return c.JSON(http.StatusNoContent, nil)
}
