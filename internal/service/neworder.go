package service

import (
	"context"
	"time"
)

func (s *Service) NewOrder(ctx context.Context, orderId string) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var err error

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		err = s.Repository.NewOrder(ctx, orderId)
		if err == nil {
			break
		} else if !IsConnectionError(err) {
			return err
		}
	}

	return err

}
