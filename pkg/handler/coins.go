package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"

	"github.com/mickey-mickser/simplecoin"
)

type getAllCoinsResponse struct {
	Data []crypto.CryptoCoin `json:"data"`
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	data, err := h.useCase.CryptoList.GetAll(ctx, nil)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(getAllCoinsResponse{Data: data}); err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

func (h *Handler) GetCoin(w http.ResponseWriter, r *http.Request) {
	coin := chi.URLParam(r, "coin")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	//ctx := context.Background()
	cryptoCoin := crypto.CryptoCoin{
		SymbolFrom: coin,
	}

	price, err := h.useCase.GetCoin(ctx, cryptoCoin)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]float64{
		"price": price,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

}

func (h *Handler) CoinToCoin(w http.ResponseWriter, r *http.Request) {

	coin := chi.URLParam(r, "coin")
	coin2 := chi.URLParam(r, "coin2")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	cryptoCoin := crypto.CryptoCoin{
		SymbolFrom: coin,
	}
	cryptoCoin2 := crypto.CryptoCoin{
		SymbolFrom: coin2,
	}

	price, err := h.useCase.GetCoin(ctx, cryptoCoin)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	price2, err := h.useCase.GetCoin(ctx, cryptoCoin2)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if price2 == 0 {
		h.newErrorResponse(w, http.StatusBadRequest, "Price of the second coin is zero")
		return
	}

	result := price / price2

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]float64{
		coin + "_to_" + coin2: result,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

}
