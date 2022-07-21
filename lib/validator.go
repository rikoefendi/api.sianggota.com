package lib

import (
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
}

func NewValidator() (cv *CustomValidator) {
	return new(CustomValidator)
}
func (cv *CustomValidator) Validate(i interface{}) error {
	v := validate.Struct(i)
	if ok := v.Validate(); !ok {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, v.Errors.All())
	}
	return nil
}
