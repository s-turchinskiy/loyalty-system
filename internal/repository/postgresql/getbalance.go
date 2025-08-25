package postgresql

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

func (p *PostgreSQL) GetBalance(ctx context.Context, userId string) (*models.Balance, error) {

	request, err := getRequest("get_balance_withdraws.sql")
	if err != nil {
		return nil, err
	}

	var result []models.Balance
	err = p.db.SelectContext(ctx, &result, request)

	if err != nil {
		return nil, common.WrapError(err)
	}

	if len(result) == 0 {

		return &models.Balance{Current: 0, Withdrawn: 0}, nil
	}

	return &result[0], nil

}
