package postgresql

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

func (p *PostgreSQL) GetOrders(ctx context.Context, userID string) ([]models.Order, error) {

	request, err := getRequest("get_orders.sql")
	if err != nil {
		return nil, err
	}

	var result []models.Order
	err = p.db.SelectContext(ctx, &result, request)

	if err != nil {
		return nil, common.WrapError(err)
	}

	return result, nil

}
