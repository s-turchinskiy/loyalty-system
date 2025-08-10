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
)

func init() {

	if err := logger.Initialize(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	ctx := context.Background()
	metricsHandler := handlers.NewHandler(ctx)
	errors := make(chan error)

	err := godotenv.Load("./cmd/gophermart/.env")
	if err != nil {
		logger.Log.Debugw("Error loading .env file", "error", err.Error())
	}

	if err := config.SetConfig(); err != nil {
		//logger.Log.Errorw("Get Settings error", "error", err.Error())
		log.Fatal(err)
	}

	go func() {
		err := run(metricsHandler)
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

func run(h *handlers.Handler) error {

	router := internal.Router(h)
	logger.Log.Info("Running server", zap.String("address", config.Config.Address.String()))
	return http.ListenAndServe(config.Config.Address.String(), router)

}
