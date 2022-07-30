package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var dbpool *sqlitex.Pool

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

func initDbPool() {
	poolSize := 10 // this is the number of connections in the pool
	var err error
	dbpool, err = sqlitex.Open("db/go-game.db", sqlite.SQLITE_OPEN_READWRITE, poolSize)
	if err != nil {
		panic(err)
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := initEcho()
	initDbPool()

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	conn := dbpool.Get(c.Request().Context())
	if conn == nil {
		return nil
	}
	// query all registered users
	var registeredUsers []User = make([]User, 0)
	defer dbpool.Put(conn) // put the conn back in the pool
	stmt := conn.Prep("SELECT id, handle FROM users;")
	var user User
	for {
		if hasRow, err := stmt.Step(); err != nil {
			c.Logger().Fatal(err)
		} else if !hasRow {
			break
		}
		user = User{
			Id:     int(stmt.GetInt64("id")),
			Handle: stmt.GetText("handle"),
		}
		registeredUsers = append(registeredUsers, user)
	}

	data := struct {
		Time            time.Time
		IsLoggedIn      bool
		LoggedInUser    User
		RegisteredUsers []User
	}{
		Time:            time.Now(),
		IsLoggedIn:      true,
		LoggedInUser:    user,
		RegisteredUsers: registeredUsers,
	}
	return c.Render(http.StatusOK, "index.html", data)
}
