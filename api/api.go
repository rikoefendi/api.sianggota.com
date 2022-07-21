package api

import (
	"api.sianggota.com/api/users"
	"github.com/labstack/echo/v4"
)

func New(e echo.Echo) {
	api := e.Group("api/")
	users.Routes(api)
}
