package main

import (
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

	e.GET("/", listTodosHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func listTodosHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}
