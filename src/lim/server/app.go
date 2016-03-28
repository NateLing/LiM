package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"lim/router"
	"runtime"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	tpl := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// Echo instance
	echo := echo.New()

	// Middleware
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	echo.Use(middleware.Gzip())
	echo.Use(middleware.Static("public"))

	// CORS
	echo.Use(standard.WrapMiddleware(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost"},
	}).Handler))

	echo.SetRenderer(tpl)
	echo.File("/favicon.ico", "/favicon.ico")

	// Websocket
	echo.Get("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		for {
			websocket.Message.Send(ws, "Hello, Client!")
			msg := ""
			websocket.Message.Receive(ws, &msg)
			println(msg)
		}
	})))

	router.AppRouter(echo)

	echo.Run(standard.New(":80"))
}
