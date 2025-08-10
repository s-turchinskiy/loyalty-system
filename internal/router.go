package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/s-turchinskiy/loyalty-system/internal/handlers"
)

func Router(h *handlers.Handler) chi.Router {

	router := chi.NewRouter()
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
