package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func HTTPErrorHandler() echo.HTTPErrorHandler {
	data := &RData{}
	return func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		page := ""
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		data.Title = code
		switch code {
		case 404:
			page = "404"
			break
		case 502:
			page = "502"
			break
		default:
			page = "unknow"
			break
		}
		HTML(c, page, data)
	}
}

func AppRouter(echo *echo.Echo) {
	echo.Get("/", Index())

	// Error handler
	echo.SetHTTPErrorHandler(HTTPErrorHandler())
}
