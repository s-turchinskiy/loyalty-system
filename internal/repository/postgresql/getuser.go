package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
)

func (p *PostgreSQL) GetUser(ctx context.Context, login string) (hashPassword string, err error) {

	request, err := getRequest("get_user.sql")
	if err != nil {
		return "", err
	}

	row := p.db.QueryRowContext(ctx, request, login)
	err = row.Scan(&hashPassword)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return "", servicecommon.NewErrorUserNoExist(login)
	case err != nil:
		return "", err
	}

	return hashPassword, nil

}
