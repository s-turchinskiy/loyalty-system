package service

import (
	"github.com/s-turchinskiy/loyalty-system/internal/repository"
	"sync"
	"time"
)

type Service struct {
	Repository    repository.Repository
	retryStrategy []time.Duration
	mutex         sync.Mutex
}
