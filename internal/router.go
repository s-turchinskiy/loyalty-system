package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/s-turchinskiy/loyalty-system/internal/handlers"
	"github.com/s-turchinskiy/loyalty-system/internal/middleware/logger"
)

func Router(h *handlers.Handler) chi.Router {

	router := chi.NewRouter()
	router.Use(logger.Logger)
	router.Use(h.AuthorizationMiddleware)
	//router.Use(gzip.GzipMiddleware)
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/orders", h.NewOrder)
		r.Get("/orders", h.GetOrders)
		r.Route("/balance", func(r chi.Router) {
			r.Get("/", h.GetBalance)
			r.Post("/withdraw", h.NewWithdraw)

		})
		r.Get("/withdrawals", h.GetWithdrawals)
	})

	return router

}
