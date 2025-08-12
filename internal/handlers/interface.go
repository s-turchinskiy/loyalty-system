package handlers

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/repository/postgresql"
	"github.com/s-turchinskiy/loyalty-system/internal/service"
	"log"
	"time"
)

type Handler struct {
	Service service.Updater
}

func NewHandler(ctx context.Context, addr, schemaName string) *Handler {

	p, err := postgresql.NewPostgresStorage(ctx, addr, schemaName)
	if err != nil {
		logger.Log.Debugw("Connect to database error", "error", err.Error())
		log.Fatal(err)
	}

	retryStrategy := []time.Duration{
		0,
		2 * time.Second,
		5 * time.Second}

	return &Handler{

		Service: service.New(p, retryStrategy)}
}
