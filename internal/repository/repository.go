package repository

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

type Repository interface {
	Ping(ctx context.Context) ([]byte, error)
	NewOrder(ctx context.Context, orderID string) error
	GetOrders(ctx context.Context, userID string) ([]models.Order, error)
	GetBalance(ctx context.Context, userID string) (*models.Balance, error)
	NewWithdraw(ctx context.Context, userID string, newWithdraw models.NewWithdraw) error
	GetWithdrawals(ctx context.Context, userID string) ([]models.Withdrawal, error)
}
