package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var todos []Todo = []Todo{
	{
		Title: "foo",
		Done:  false,
	},
}

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", listTodosHandler)
	e.POST("/todos", addTodoHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func listTodosHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func addTodoHandler(c echo.Context) error {
	newTodo := new(Todo)
	if err := c.Bind(newTodo); err != nil {
		fmt.Printf("Could not bind input data. %v", err.Error())
		return err
	}

	todos = append(todos, *newTodo)

	uri := fmt.Sprintf("/todos/%d", len(todos)-1)
	return c.String(http.StatusCreated, uri)
}
