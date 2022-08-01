package users

import (
	"net/http"
	"strconv"

	"api.sianggota.com/utils"
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
		return err
	}
	if err = c.Validate(input); err != nil {
		return err
	}
	result, err := h.r.Create(*input)
	if err != nil {
		return c.JSON(500, err)
	}
	r := utils.Response(result)
	return c.JSON(http.StatusOK, r)
}

func (h *UserHandler) Update(c echo.Context) (err error) {
	input := UserUpdateInput{}
	id := c.Param("id")
	if err = c.Bind(&input); err != nil {
		return err
	}
	if err = c.Validate(input); err != nil {
		return err
	}
	dest := Model{
		Name: &input.Name,
	}
	result, err := h.r.UpdateById(id, dest)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utils.Response(result))
}

func (h *UserHandler) Show(c echo.Context) (err error) {
	id := c.Param("id")
	user, err := h.r.FetchById(id)
	if err != nil {
		return err
	}
	code := http.StatusOK
	return c.JSON(code, utils.Response(user))
}

func (h *UserHandler) Index(c echo.Context) (err error) {
	perPage, _ := strconv.Atoi(c.QueryParam("per_page"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	users, paginated, err := h.r.ShowAll(page, perPage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, utils.Response(users).SetMeta(paginated))
}
