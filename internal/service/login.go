package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
	"time"
)

func (s *Service) Login(ctx context.Context, user models.User) (hashPassword string, err error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		hashPassword, err = s.Repository.GetUser(ctx, user.Login)
		if err == nil {
			break
		} else if !IsConnectionError(err) {

			return "", err
		}
	}

	if !common.HashAndPasswordEqual(user.Password, hashPassword) {
		return "", servicecommon.NewErrorLoginPasswordIncorrect(user.Login)
	}
	return hashPassword, err

}
