package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// func IdAuth(eh echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		idParam := c.Param("id")
// 		id, err := strconv.ParseFloat(idParam, 64)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]any{
// 				"messege": "invalid id",
// 			})
// 		}

// 		token := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)

// 		userId := token["id"].(float64)
// 		if userId == id {
// 			return eh(c)
// 		}

// 		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})

// 	}
// }

func IdAuth(c echo.Context) float64 {

	token := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)

	userId := token["id"].(float64)
	return userId

}
