package users

import (
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	r Repository
}

func Handler(r Repository) (h *UserHandler) {
	h = &UserHandler{r}
	return h
}

func (h *UserHandler) Create(c echo.Context) (err error) {
	input := &UserCreateInput{}
	if err = c.Bind(input); err != nil {
		return c.JSON(400, err.Error())
	}
	result, err := h.r.Create(*input)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, result)
}
