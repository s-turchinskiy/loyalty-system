package service

import (
	"github.com/s-turchinskiy/loyalty-system/internal/repository"
	"time"
)

type Updater interface {
}

func New(rep repository.Repository, retryStrategy []time.Duration) *Service {

	if len(retryStrategy) == 0 {
		retryStrategy = []time.Duration{0}
	}
	return &Service{
		Repository:    rep,
		retryStrategy: retryStrategy,
	}
}
