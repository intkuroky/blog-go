package errors

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"strings"
)

var Debug bool

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if (strings.Compare(err.Error(), "code=400, message=Missing or malformed jwt") == 0) {
		c.Redirect(http.StatusMovedPermanently,"/users/login")
	}

	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else if Debug {
		msg = err.Error()
		if he.Inner != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Inner)
		}
	} else {
		msg = http.StatusText(code)
	}

	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, err.Error())
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}

	c.Logger().Error(err)
}