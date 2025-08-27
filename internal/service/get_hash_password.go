package service

import (
	"context"
	"time"
)

func (s *Service) GetHashedPassword(ctx context.Context, login string) (hashPassword string, err error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, delay := range s.retryStrategy {
		time.Sleep(delay)
		hashPassword, err = s.Repository.GetUser(ctx, login)
		if err == nil {
			break
		} else if !IsConnectionError(err) {

			return "", err
		}
	}

	return hashPassword, err

}
