package repository

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

type Repository interface {
	Ping(ctx context.Context) ([]byte, error)
	NewUser(ctx context.Context, login, hash string) error
	GetUser(ctx context.Context, login string) (hash string, err error)
	GetOrderAlreadyUploaded(ctx context.Context, orderID string) (uploaded bool, who string, err error)
	NewOrder(ctx context.Context, login, orderID string) error
	GetOrders(ctx context.Context, login string) ([]models.Order, error)
	GetBalance(ctx context.Context, login string) (*models.Balance, error)
	NewWithdraw(ctx context.Context, login string, newWithdraw models.NewWithdraw) error
	GetWithdrawals(ctx context.Context, login string) ([]models.Withdrawal, error)
}
