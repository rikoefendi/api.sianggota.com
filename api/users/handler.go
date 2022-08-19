package users

import (
	"net/http"
	"strconv"

	"api.sianggota.com/lib"
	"api.sianggota.com/utils"
	"github.com/gookit/event"
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
	if err = lib.BindValidate(c, input); err != nil {
		return err
	}
	result, err := h.r.Create(*input)
	if err != nil {
		return err
	}
	event.Fire("users.create", event.M{"users": result})
	r := utils.Response(result)
	return c.JSON(http.StatusOK, r)
}

func (h *UserHandler) Update(c echo.Context) (err error) {
	input := UserUpdateInput{}
	id := c.Param("id")
	if err = lib.BindValidate(c, &input); err != nil {
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

func (h *UserHandler) Destroy(c echo.Context) (err error) {
	id := c.Param("id")
	permanent := c.QueryParam("permanent")
	db := h.r.db
	if permanent == "true" {
		db = db.Unscoped()
	}
	user := &Model{}
	result := db.Where("id = ?", id).Delete(&user)
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, utils.Response(nil).SetMessage("successfully"))
}
