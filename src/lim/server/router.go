package main

import (
	"github.com/labstack/echo"
	"lim/router"
	"lim/router/chat"
)

func AppRouter(e *echo.Echo) {
	e.Get("/", router.Index())
	e.Get("/im", chat.Im())
}
