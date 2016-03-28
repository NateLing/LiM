package router

import (
	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	data := &RData{
		Title: "Index",
	}
	return func(c echo.Context) error {
		return HTML(c, "index", data)
	}
}
