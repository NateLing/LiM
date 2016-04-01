package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"golang.org/x/net/websocket"
	"lim/router"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Echo instance
	echo := echo.New()

	// Middleware
	echo.Use(middleware.LoggerFromConfig(middleware.LoggerConfig{
		Format: "method=${method}, status=${status}, uri=${uri},  took=${response_time}, sent=${response_size} bytes\n",
	}))
	echo.Use(middleware.Recover())
	/*echo.Use(middleware.GzipFromConfig(middleware.GzipConfig{
		Level: -1,
	}))*/
	echo.Use(middleware.Static("public"))

	// CORS
	echo.Use(standard.WrapMiddleware(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost"},
	}).Handler))

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
