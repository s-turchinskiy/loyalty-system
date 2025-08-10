package main

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal"
	"github.com/s-turchinskiy/loyalty-system/internal/handlers"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	metricsHandler := handlers.NewHandler(ctx)
	errors := make(chan error)

	go func() {
		err := run(metricsHandler)
		if err != nil {

			//logger.Log.Errorw("Server startup error", "error", err.Error())
			errors <- err
			return
		}
	}()

	err := <-errors
	//metricsHandler.Service.SaveMetricsToFile(ctx)
	//logger.Log.Infow("error, server stopped", "error", err.Error())
	log.Fatal(err)

}

func run(h *handlers.Handler) error {

	router := internal.Router(h)

	//logger.Log.Info("Running server", zap.String("address", settings.Settings.Address.String()))

	return http.ListenAndServe(":8080", router)

}
