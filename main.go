package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/", hello)
	return e
}
func main() {
	e := initEcho()
	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
