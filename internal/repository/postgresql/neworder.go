package postgresql

import (
	"context"
)

func (p *PostgreSQL) NewOrder(ctx context.Context, orderID string) error {

	request, err := getRequest("create_order.sql")
	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, request, orderID)
	return err

}
