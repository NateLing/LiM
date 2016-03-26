package router

import (
	"github.com/labstack/echo"
)

func AppRouter(r *echo.Echo) {
	r.Get("/", Index())
}
