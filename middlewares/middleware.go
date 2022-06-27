package middlewares

import "github.com/labstack/echo/v4"

func New(c *echo.Echo) {
	RequestIDMiddleware(c)
}
