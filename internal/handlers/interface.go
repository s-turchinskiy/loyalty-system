package handlers

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/repository/postgresql"
	"github.com/s-turchinskiy/loyalty-system/internal/service"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

const (
	ContentTypeTextHTML         = "text/html; charset=utf-8"
	ContentTypeTextPlain        = "text/plain"
	ContentTypeTextPlainCharset = "text/plain; charset=utf-8"
	ContentTypeApplicationJSON  = "application/json"

	TextErrorGettingData = "error getting data"
)

type Handler struct {
	Service  service.Updater
	tokenExp time.Duration
}

func NewHandler(ctx context.Context, addr, schemaName string, tokenExp time.Duration) *Handler {

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
		Service:  service.New(p, retryStrategy),
		tokenExp: tokenExp,
	}
}

func errorGettingData(w http.ResponseWriter, err error) {
	logger.Log.Infow(TextErrorGettingData, zap.Error(err))
	w.Header().Set("Content-Type", ContentTypeTextPlainCharset)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func internalError(w http.ResponseWriter, err error) {
	logger.Log.Warnw("internal error", zap.Error(err))
	w.Header().Set("Content-Type", ContentTypeTextPlainCharset)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
