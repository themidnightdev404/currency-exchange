package handlers

import (
	"currency-exchange/internal/repositories"
	"encoding/json"
	"net/http"
)

type CurrencyHandler struct {
	repo *repositories.CurrencyRepository
}

func NewCurrencyHandler(repo *repositories.CurrencyRepository) *CurrencyHandler {
	return &CurrencyHandler{repo: repo}
}

func (h *CurrencyHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (h *CurrencyHandler) GetCurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.FindAllCurrencies()
	if err != nil {
		http.Error(w, "база данных недоступна", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
