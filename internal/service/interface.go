package service

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/repository"
	"sync"
	"time"
)

type Updater interface {
	Register(ctx context.Context, newUser models.NewUser) (hashPassword string, err error)
	NewOrder(ctx context.Context, orderID string) error
	GetOrders(ctx context.Context, userID string) (models.Orders, error)
	GetBalance(ctx context.Context, userID string) (*models.Balance, error)
	NewWithdraw(ctx context.Context, userID string, newWithdraw models.NewWithdraw) error
	GetWithdrawals(ctx context.Context, userID string) (models.Withdrawals, error)
}

type Service struct {
	Repository    repository.Repository
	retryStrategy []time.Duration
	mutex         sync.Mutex
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

func IsConnectionError(err error) bool {

	if err == nil {
		return false
	}

	var pgErr *pgconn.PgError
	errors.As(err, &pgErr)
	if pgErr == nil {
		return false
	}

	return pgerrcode.IsConnectionException(pgErr.Code)

}
