package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
)

func (s *Service) Login(ctx context.Context, user models.User) (hashPassword string, err error) {

	hashPassword, err = s.GetHashedPassword(ctx, user.Login)

	if !common.HashAndPasswordEqual(user.Password, hashPassword) {
		return "", servicecommon.NewErrorLoginPasswordIncorrect(user.Login)
	}
	return hashPassword, err

}
