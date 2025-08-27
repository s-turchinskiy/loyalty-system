package authservice

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"time"
)

func BuildJWTString(login, hashPassword string, tokenExp time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte(hashPassword))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
