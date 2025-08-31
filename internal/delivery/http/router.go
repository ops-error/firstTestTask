package http

import (
	"firstTestTask/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(repo *repository.OrderRepo) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(res http.ResponseWriter, _ *http.Request) {
		res.Write([]byte("Hello World"))
	})
	router.Get("/orders/{uid}", GetFullOrder(repo))
	return router
}
