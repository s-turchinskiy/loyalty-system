package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/s-turchinskiy/loyalty-system/internal/common"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/repository"
	"time"
)

type PostgreSQL struct {
	db   *sqlx.DB
	pool *pgxpool.Pool
}

func NewPostgresStorage(ctx context.Context, addr, schemaName string) (repository.Repository, error) {

	logger.Log.Debug("addr for Sql.Open: ", addr)

	db, err := sqlx.Open("pgx", addr)
	if err != nil {
		return nil, common.WrapError(err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, common.WrapError(err)
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	pool, err := pgxpool.New(ctx, addr)
	if err != nil {
		return nil, common.WrapError(err)
	}

	p := &PostgreSQL{db: db, pool: pool}

	_, err = p.db.ExecContext(ctx, fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s`, schemaName))
	if err != nil {
		return nil, common.WrapError(err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{SchemaName: schemaName})
	if err != nil {
		return nil, common.WrapError(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, common.WrapError(err)
	}

	return p, nil

}
