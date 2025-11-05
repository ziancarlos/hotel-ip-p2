package middleware

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if customErr, ok := err.(*exception.CustomError); ok {
		c.JSON(customErr.Code, web.WebResponse{
			Message: customErr.Message,
		})
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, web.WebResponse{
			Message: he.Message.(string),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, web.WebResponse{
		Message: "Internal server error",
	})
}
