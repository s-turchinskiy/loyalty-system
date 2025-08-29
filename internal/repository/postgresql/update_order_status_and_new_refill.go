package postgresql

import (
	"context"
)

func (p *PostgreSQL) UpdateOrderStatusAndNewRefill(ctx context.Context, orderID, newStatus string, sum float64, userID uint) error {

	request1, err := getRequest("update_order_status.sql")
	if err != nil {
		return err
	}

	request2, err := getRequest("create_refill.sql")
	if err != nil {
		return err
	}

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, request1, orderID, newStatus)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, request2, orderID, sum, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
