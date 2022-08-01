package lib

import (
	"net/http"

	"api.sianggota.com/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	code int
	res  *utils.ResponseStr
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	res = utils.Response(nil)
	switch err {
	case gorm.ErrRecordNotFound:
		code = http.StatusNotFound
		res.SetErrors(err.Error()).SetMessage(http.StatusText(code))
	default:
		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}
		// ohje, ov := err.(utils.ResponseStr)
		// Issue #1426
		code = he.Code
		r, ok := he.Message.(*utils.ResponseStr)
		res = utils.Response(nil).SetErrors(he).SetMessage(http.StatusText(code))
		if ok {
			res = r
		}
	}

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, res)
	}
	if err != nil {
		c.Logger().Error(err)
	}
	// c.JSON(code, utils.Response(nil).SetErrors(err).SetMessage(http.StatusText(code)))
}
