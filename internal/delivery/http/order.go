package http

import (
	"encoding/json"
	"firstTestTask/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetFullOrder(repo *repository.OrderRepo) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		uid := chi.URLParam(req, "uid")
		order, err := repo.GetFullOrder(req.Context(), uid)
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(res).Encode(order)
	}
}
