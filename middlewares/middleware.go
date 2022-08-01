package middlewares

import (
	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) {
	RequestIDMiddleware(e)
	e.Use(Logger())
}
