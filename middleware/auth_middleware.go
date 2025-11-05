package middleware

import (
	"hotel_ip-p2/exception"
	"hotel_ip-p2/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return exception.NewCustomError(http.StatusUnauthorized, "Missing authorization header")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return exception.NewCustomError(http.StatusUnauthorized, "Invalid authorization header format")
		}

		token := tokenParts[1]
		claims, err := helper.ValidateToken(token)
		if err != nil {
			return exception.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
		}

		c.Set("user_id", claims.UserID)

		return next(c)
	}
}
