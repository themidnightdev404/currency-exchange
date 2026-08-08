package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/themidnightdev404/currency-exchange/internal/models"
	"github.com/themidnightdev404/currency-exchange/internal/service"
)

type CurrencyHandler struct {
	service *service.CurrencyService
}

func NewCurrencyHandler(service *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}

func (h *CurrencyHandler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /currencies", h.GetCurrencies)
	mux.HandleFunc("POST /currencies", h.CreateCurrency)
	mux.HandleFunc("GET /currency/{code}", h.GetCurrencyByCode)

	return mux
}

func (h *CurrencyHandler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.service.GetCurrencies(r.Context())
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currencies)
}

func (h *CurrencyHandler) GetCurrencyByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if len(code) != 3 {
		http.Error(w, "Неверный формат кода", http.StatusBadRequest)
		return
	}

	currency, err := h.service.GetCurrencyByCode(r.Context(), code)
	if err != nil {
		http.Error(w, "Валюта не найдена в базе", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currency)
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var currency models.Currency

	err := json.NewDecoder(r.Body).Decode(&currency)
	if err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	createdCurrency, err := h.service.CreateCurrency(r.Context(), &currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(createdCurrency)
	if err != nil {
		return
	}
}
