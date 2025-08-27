package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/s-turchinskiy/loyalty-system/cmd/gophermart/config"
	"github.com/s-turchinskiy/loyalty-system/internal"
	"github.com/s-turchinskiy/loyalty-system/internal/handlers"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

func init() {

	if err := logger.Initialize(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	err := godotenv.Load("./cmd/gophermart/.env")
	if err != nil {
		logger.Log.Debugw("Error loading .env file", "error", err.Error())
	}

	config, err := config.GetConfig()
	if err != nil {
		logger.Log.Errorw("Get Settings error", "error", err.Error())
		log.Fatal(err)
	}

	ctx := context.Background()
	handler := handlers.NewHandler(
		ctx,
		config.Database.String(),
		config.Database.DBName,
		time.Hour*1000)
	errors := make(chan error)

	go func() {
		err := run(handler, config.Address.String())
		if err != nil {

			logger.Log.Errorw("Server startup error", "error", err.Error())
			errors <- err
			return
		}
	}()

	err = <-errors
	logger.Log.Infow("error, server stopped", "error", err.Error())
	log.Fatal(err)

}

func run(h *handlers.Handler, addr string) error {

	router := internal.Router(h)
	logger.Log.Info("Running server", zap.String("address", addr))
	return http.ListenAndServe(addr, router)

}
