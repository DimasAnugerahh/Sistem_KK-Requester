package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authorization(c echo.Context) (float64, string) {

	token := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)

	userId := token["id"].(float64)
	userRole := token["role"].(string)

	return userId, userRole

}
