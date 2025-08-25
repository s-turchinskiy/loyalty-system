package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"time"
)

func (s *Service) NewWithdraw(ctx context.Context, userId string, newWithdraw models.NewWithdraw) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var err error

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		err = s.Repository.NewWithdraw(ctx, userId, newWithdraw)
		if err == nil {
			break
		} else if !IsConnectionError(err) {
			return err
		}
	}

	return err

}
