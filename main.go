package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initEcho() *echo.Echo {
	e := echo.New()
	// pre-compile the templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t // register the templates with Echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/", index)
	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := initEcho()

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	data := struct {
		Time            time.Time
		IsLoggedIn      bool
		LoggedInUser    User
		RegisteredUsers []User
	}{
		Time:       time.Now(),
		IsLoggedIn: false,
	}
	return c.Render(http.StatusOK, "index.html", data)
}
