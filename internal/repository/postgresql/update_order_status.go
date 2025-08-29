package postgresql

import (
	"context"
)

func (p *PostgreSQL) UpdateOrderStatus(ctx context.Context, orderID, newStatus string) error {

	request, err := getRequest("update_order_status.sql")
	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, request, orderID, newStatus)
	return err

}
