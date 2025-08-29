package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/s-turchinskiy/loyalty-system/cmd/gophermart/config"
	"github.com/s-turchinskiy/loyalty-system/internal"
	"github.com/s-turchinskiy/loyalty-system/internal/accrualservice"
	"github.com/s-turchinskiy/loyalty-system/internal/handlers"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"github.com/s-turchinskiy/loyalty-system/internal/repository/postgresql"
	"go.uber.org/zap"
	"log"
	"net/http"
	"runtime"
	"time"
)

func init() {

	if err := logger.Initialize(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	ctx := context.Background()

	err := godotenv.Load("./cmd/gophermart/.env")
	if err != nil {
		logger.Log.Debugw("Error loading .env file", "error", err.Error())
	}

	config, err := config.GetConfig()
	if err != nil {
		logger.Log.Errorw("Get Settings error", "error", err.Error())
		log.Fatal(err)
	}

	repository, err := postgresql.NewPostgresStorage(ctx, config.Database.String(), config.Database.DBName)
	if err != nil {
		logger.Log.Debugw("Connect to database error", "error", err.Error())
		log.Fatal(err)
	}

	retryStrategy := []time.Duration{
		0,
		2 * time.Second,
		5 * time.Second}

	handler := handlers.NewHandler(
		ctx,
		repository,
		retryStrategy,
		time.Hour*1000)

	errorsCh := make(chan error)

	go func() {
		err := run(handler, config.Address.String())
		if err != nil {

			logger.Log.Errorw("Server startup error", "error", err.Error())
			errorsCh <- err
			return
		}
	}()

	doneCh := make(chan struct{})
	defer close(doneCh)

	accrualService := accrualservice.New(
		ctx,
		repository,
		retryStrategy,
		time.Duration(2)*time.Second,
		accrualservice.NewHTTPResty(fmt.Sprintf("%s/api/orders/{number}", config.AccrualSystem)),
		runtime.NumCPU(),
		errorsCh,
		doneCh)

	go accrualService.RunPeriodically()

	err = <-errorsCh
	logger.Log.Infow("error, server stopped", "error", err.Error())
	log.Fatal(err)

}

func run(h *handlers.Handler, addr string) error {

	router := internal.Router(h)
	logger.Log.Info("Running server", zap.String("address", addr))
	return http.ListenAndServe(addr, router)

}
