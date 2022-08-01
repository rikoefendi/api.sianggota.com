package users

import (
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	repo := New()
	handler := Handler(*repo)
	users := g.Group("users")
	users.POST("", handler.Create)
	users.PUT("/:id", handler.Update)
	users.GET("/:id", handler.Show)
	users.GET("", handler.Index)
	users.DELETE("/:id", handler.Destroy)
}
