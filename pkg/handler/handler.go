package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/mickey-mickser/simplecoin/pkg/usecase"
)

type Handler struct {
	useCase *usecase.UseCase
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//

	router.Route("/api", func(r chi.Router) {
		r.Get("/coins", h.GetAll)
		r.Get("/{coin}", h.GetCoin)
		r.Get("/{coin}/{coin2}", h.CoinToCoin)
	})
	return router
}
