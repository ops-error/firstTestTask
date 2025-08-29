package http

import (
	"firstTestTask/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewRouter(dataBase *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	repo := repository.NewOrderRepo(dataBase)
	router.Get("/", func(res http.ResponseWriter, _ *http.Request) {
		res.Write([]byte("Hello World"))
	})
	router.Get("/orders/{uid}", GetFullOrder(repo))
	return router
}
