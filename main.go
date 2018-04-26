package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"rename/rename"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port string
)

func indexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	// Parse the args
	flag.StringVar(&port, "port", ":80", "Set the port for the application to listen to, else it will default to port 80")

}

func main() {

	flag.Parse()
	log.Println(port)
	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	e := echo.New()
	e.Use(middleware.Static("public"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	// Routes
	e.GET("/", indexHandler)

	// API
	g := e.Group("/api")
	{
		g.GET("/tables", rename.APITablesHandler)
	}

	rename.InitDB()
	e.Logger.SetOutput(os.Stdout)
	// Lets start this bad boi
	e.Logger.Fatal(e.Start(port))

}
