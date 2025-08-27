package postgresql

import (
	"context"
	"database/sql"
	"errors"
)

func (p *PostgreSQL) GetOrderAlreadyUploaded(ctx context.Context, orderID string) (uploaded bool, who string, err error) {

	request, err := getRequest("get_order_already_uploaded.sql")
	if err != nil {
		return false, "", err
	}

	row := p.db.QueryRowContext(ctx, request, orderID)
	err = row.Scan(&who)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return false, "", nil
	case err != nil:
		return false, "", err
	}

	return true, who, nil

}
