package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/s-turchinskiy/loyalty-system/internal/models"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
)

func (p *PostgreSQL) NewWithdraw(ctx context.Context, login string, newWithdraw models.NewWithdraw) error {

	requestSelect, err := getRequest("get_balance.sql")
	if err != nil {
		return err
	}

	requestInsert, err := getRequest("create_withdraw.sql")
	if err != nil {
		return err
	}

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	var balance float64
	row := tx.QueryRowContext(ctx, requestSelect, login)
	err = row.Scan(&balance)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		tx.Rollback()
		return servicecommon.ErrorNotEnoughBalance
	case err != nil:
		tx.Rollback()
		return err
	}

	if balance < newWithdraw.Sum {
		return servicecommon.ErrorNotEnoughBalance
	}

	_, err = tx.Exec(requestInsert, newWithdraw.Order, -1*newWithdraw.Sum, login)
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
