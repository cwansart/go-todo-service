package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
