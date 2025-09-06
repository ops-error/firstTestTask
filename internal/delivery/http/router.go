package http

import (
	"firstTestTask/internal/repository"

	"github.com/go-chi/chi/v5"
)

func NewRouter(repo *repository.OrderRepo) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/orders/{uid}", GetFullOrder(repo))
	return router
}
