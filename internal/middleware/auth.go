package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/helper/jwt"
)

func Auth(next echo.HandlerFunc, roles ...string) echo.HandlerFunc {
	return func(e echo.Context) error {
		Authorization := e.Request().Header.Get("Authorization")

		if len(Authorization) < 9 || Authorization[:7] != "Bearer " {
			return e.JSON(http.StatusUnauthorized, exception.NewUnauthorized("Invalid token"))
		}

		token := Authorization[7:]
		claim, err := jwt.ClaimToken(token)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, exception.NewUnauthorized("Invalid token"))
		}

		for _, role := range roles {
			if role == claim.Role {
				e.Set("user", claim)
				return next(e)
			}
		}
		return e.JSON(http.StatusUnauthorized, exception.NewUnauthorized("You don't have access"))

	}
}
