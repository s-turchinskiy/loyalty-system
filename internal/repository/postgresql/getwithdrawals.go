package postgresql

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
)

func (p *PostgreSQL) GetWithdrawals(ctx context.Context, userId string) ([]models.Withdrawal, error) {

	request, err := getRequest("get_withdrawals.sql")
	if err != nil {
		return nil, err
	}

	var result []models.Withdrawal
	err = p.db.SelectContext(ctx, &result, request)

	if err != nil {
		return nil, common.WrapError(err)
	}

	return result, nil

}
