package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestIDMiddleware(e *echo.Echo) {
	e.Use(middleware.RequestID())
}
