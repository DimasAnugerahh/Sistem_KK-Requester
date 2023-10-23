package helper

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

func RoleAuth(eh echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)

		if userRole, roleExists := token["role"].(string); roleExists {
			LowUserRole := strings.ToLower(userRole)
			if LowUserRole == "admin" {
				return eh(c)
			}
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})

	}
}
