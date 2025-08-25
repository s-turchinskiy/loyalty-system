package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"time"
)

func (s *Service) GetBalance(ctx context.Context, userID string) (*models.Balance, error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var err error

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		result, err := s.Repository.GetBalance(ctx, "")
		if err == nil {
			return result, nil
		} else if !IsConnectionError(err) {
			return nil, err
		}
	}

	return nil, err

}
