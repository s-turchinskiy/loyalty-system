package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"time"
)

func (s *Service) Register(ctx context.Context, newUser models.User) (hashPassword string, err error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	hash, err := common.HashFromPassword(newUser.Password)
	if err != nil {
		return "", err
	}

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		err = s.Repository.NewUser(ctx, newUser.Login, hash)
		if err == nil {
			break
		} else if !IsConnectionError(err) {
			return "", err
		}
	}

	return hash, err

}
