package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Echo instance
	mux := echo.New()
	mux.Static("/", "public")

	// Middleware
	mux.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, status=${status}, uri=${uri}\n",
	}))
	mux.Use(middleware.Recover())
	mux.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// CORS
	mux.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
	}))

	// Websocket
	mux.Get("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		for {
			websocket.Message.Send(ws, "Hello, Client!")
			msg := ""
			websocket.Message.Receive(ws, &msg)
			println(msg)
		}
	})))

	AppRouter(mux)

	mux.Run(standard.New(":80"))
}
