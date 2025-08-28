package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(res http.ResponseWriter, _ *http.Request) {
		res.Write([]byte("Hello World"))
	})
	return router
}
