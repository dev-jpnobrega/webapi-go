package infrastructure

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func addMiddleware(server *echo.Echo) {
	server.Use(mid.Recover())
}

func CreateHTTPServer() (server *echo.Echo) {
	server = echo.New()
	server.Binder = &CustomBinder{}
	server.Validator = &CustomValidator{validator: validator.New()}

	addMiddleware(server)

	return server
}
