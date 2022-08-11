package main

import (
	"errors"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := initEcho()
	initDbPool()

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

func initEcho() *echo.Echo {
	e := echo.New()
	e = initTemplates(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/games/new", gamesNew)
	e.GET("/games", games)
	e.GET("/", index)
	return e
}

func initTemplates(e *echo.Echo) *echo.Echo {
	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("view/index.html", "view/base.html"))
	templates["games_new.html"] = template.Must(template.ParseFiles("view/games_new.html", "view/base.html"))
	templates["games.html"] = template.Must(template.ParseFiles("view/games.html", "view/base.html"))
	e.Renderer = &TemplateRegistry{templates}
	return e
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

type Data struct {
	Header          string
	IsLoggedIn      bool
	LoggedInUser    *User
	RegisteredUsers *[]User
	Games           *[]Game
}

func index(c echo.Context) error {
	// query all registered users
	//registeredUsers, err := selectAllUsers(c.Request().Context())
	//if err != nil {
	//	return err
	//}
	//user := registeredUsers[0]
	data := Data{
		Header:          "written in Go",
		IsLoggedIn:      false,
		LoggedInUser:    nil,
		RegisteredUsers: nil,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func gamesNew(c echo.Context) error {
	data := Data{
		IsLoggedIn:      false,
		LoggedInUser:    nil,
		RegisteredUsers: nil,
		Header:          "New Game",
	}
	return c.Render(http.StatusOK, "games_new.html", data)
}

func games(c echo.Context) error {
	games, err := selectAllGames(c.Request().Context())
	if err != nil {
		return err
	}
	data := Data{
		IsLoggedIn: false,
		Header:     "New Game",
		Games:      &games,
	}
	return c.Render(http.StatusOK, "games.html", data)
}
