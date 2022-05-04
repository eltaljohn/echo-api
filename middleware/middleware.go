package middleware

import (
	"github.com/eltaljohn/echo-api/authorization"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Does not has authorization"})
		}

		return f(c)
	}
}
