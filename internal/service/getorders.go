package service

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"time"
)

func (s *Service) GetOrders(ctx context.Context, userID string) (models.Orders, error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var err error

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		orders, err := s.Repository.GetOrders(ctx, userID)
		if err == nil {
			return orders, nil
		} else if !IsConnectionError(err) {
			return nil, err
		}
	}

	return nil, err

}
