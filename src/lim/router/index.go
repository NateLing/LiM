package router

import (
	"github.com/labstack/echo"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return HTML(c, "index", nil)
	}
}
