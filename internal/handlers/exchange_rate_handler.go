package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/themidnightdev404/currency-exchange/internal/service"
)

type ExchangeRateHandler struct {
	service *service.ExchangeRateService
}

func NewExchangeRateHandler(service *service.ExchangeRateService) *ExchangeRateHandler {
	return &ExchangeRateHandler{service: service}
}

func (h *ExchangeRateHandler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /exchangeRates", h.GetRates)

	return mux
}

func (h *ExchangeRateHandler) GetRates(w http.ResponseWriter, r *http.Request) {
	exchangeRates, err := h.service.GetRates(r.Context())
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchangeRates)
}
