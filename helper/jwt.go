package helper

import (
	"kk-requester/model/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userLoginResponse *domain.Account, id int, role string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return validToken, nil
}
