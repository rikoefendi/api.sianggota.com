package users

import (
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	if err = c.Validate(input); err != nil {
		return err
	}
	result, err := h.r.Create(*input)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, result)
}

func (h *UserHandler) Update(c echo.Context) (err error) {
	input := UserUpdateInput{}
	id := c.Param("id")
	if err = c.Bind(&input); err != nil {
		return c.JSON(400, err.Error())
	}
	if err = c.Validate(input); err != nil {
		return err
	}
	dest := Model{
		Name: &input.Name,
	}
	result, err := h.r.UpdateById(id, dest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, err.Error())
		}
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, result)
}
