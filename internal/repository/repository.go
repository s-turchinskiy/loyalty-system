package repository

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

type Repository interface {
	Ping(ctx context.Context) ([]byte, error)
	NewOrder(ctx context.Context, orderId string) error
	GetOrders(ctx context.Context, userId string) ([]models.Order, error)
	GetBalance(ctx context.Context, userId string) (*models.Balance, error)
	NewWithdraw(ctx context.Context, userId string, newWithdraw models.NewWithdraw) error
	GetWithdrawals(ctx context.Context, userId string) ([]models.Withdrawal, error)
}
