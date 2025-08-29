package postgresql

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

func (p *PostgreSQL) GetOrdersForAccrualCalculation(ctx context.Context) (result []models.OrdersForAccrualCalculation, err error) {

	request, err := getRequest("get_orders_for_accrual_calculation.sql")
	if err != nil {
		return nil, err
	}

	err = p.db.SelectContext(ctx, &result, request)

	if err != nil {
		return nil, common.WrapError(err)
	}

	return result, nil

}
