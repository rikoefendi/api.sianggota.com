package lib

import (
	"fmt"
	"net/http"

	"api.sianggota.com/database"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
}

func NewValidator() (cv *CustomValidator) {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
	validate.AddValidator("unique", func(value string, table string, field string) bool {
		return cv.Query(value, table, field)
	})
	validate.AddValidator("exists", func(value string, table string, field string) bool {
		return cv.Query(value, table, field)
	})
	validate.AddGlobalMessages(map[string]string{
		"unique": "{field} already exist",
		"exists": "{field} not exist",
	})
	return new(CustomValidator)
}
func (cv *CustomValidator) Validate(i interface{}) error {
	v := validate.Struct(i)
	if ok := v.Validate(); !ok {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, v.Errors.All())
	}
	return nil
}

func (cb *CustomValidator) Query(value string, table string, field string) bool {
	var count int64
	db, err := database.Session().DB()
	if err != nil {
		panic(err)
	}
	res := db.QueryRow(fmt.Sprintf(`SELECT count("%s") FROM %s WHERE %s='%s'`, field, table, field, value))
	if res.Err() != nil {
		panic(res.Err())
	}
	res.Scan(&count)
	return count < 1
}
