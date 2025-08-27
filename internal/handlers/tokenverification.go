package handlers

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

func (h *Handler) tokenVerification(ctx context.Context, tokenString string) (success bool, login string, err error) {

	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			hashedPassword, err := h.Service.GetHashedPassword(ctx, claims.Login)
			if err != nil {
				return nil, err
			}

			return []byte(hashedPassword), nil
		})
	if err != nil {
		return false, "", err
	}

	if !token.Valid {
		return false, "", err
	}

	return true, claims.Login, nil
}
