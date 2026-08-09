package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/themidnightdev404/currency-exchange/internal/models"
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
	mux.HandleFunc("GET /exchangeRate/{pair}", h.GetRate)
	mux.HandleFunc("POST /exchangeRates", h.CreateRate)
	mux.HandleFunc("PATCH /exchangeRate/{pair}", h.UpdateRate)

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

func (h *ExchangeRateHandler) GetRate(w http.ResponseWriter, r *http.Request) {
	pair := r.PathValue("pair")
	if len(pair) != 6 {
		http.Error(w, "Неверный формат валютной пары", http.StatusBadRequest)
		return
	}
	baseCode := pair[:3]
	targetCode := pair[3:]
	rates, err := h.service.GetRates(r.Context())
	if err != nil {
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	var foundRate *models.ExchangeRate
	for i, r := range rates {
		if r.BaseCurrency.Code == baseCode && r.TargetCurrency.Code == targetCode {
			foundRate = &rates[i]
			break
		}
	}
	if foundRate == nil {
		http.Error(w, "Курс для указанной пары не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundRate)
}

func (h *ExchangeRateHandler) CreateRate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Не удалось распарсить форму", http.StatusBadRequest)
		return
	}
	baseCurrencyIdStr := r.PostFormValue("baseCurrencyId")
	targetCurrencyIdStr := r.PostFormValue("targetCurrencyId")
	rateStr := r.PostFormValue("rate")
	if baseCurrencyIdStr == "" || targetCurrencyIdStr == "" || rateStr == "" {
		http.Error(w, "Отсутствуют обязательные параметры формы", http.StatusBadRequest)
		return
	}
	baseID, err := strconv.ParseInt(baseCurrencyIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Неверный формат ID базовой валюты", http.StatusBadRequest)
		return
	}
	targetID, err := strconv.ParseInt(targetCurrencyIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Неверный формат ID целевой валюты", http.StatusBadRequest)
		return
	}
	rateValue, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат курса", http.StatusBadRequest)
		return
	}
	exchangeRate := models.ExchangeRate{
		BaseCurrency: models.Currency{
			ID: baseID,
		},
		TargetCurrency: models.Currency{
			ID: targetID,
		},
		Rate: rateValue,
	}
	createdRate, err := h.service.CreateRate(r.Context(), &exchangeRate)
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdRate)
}

func (h *ExchangeRateHandler) UpdateRate(w http.ResponseWriter, r *http.Request) {
	pair := r.PathValue("pair")
	if len(pair) != 6 {
		http.Error(w, "Неверный формат валютной пары", http.StatusBadRequest)
		return
	}
	baseCode := pair[:3]
	targetCode := pair[3:]
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Не удалось распарсить форму", http.StatusBadRequest)
		return
	}
	rateStr := r.PostFormValue("rate")
	if rateStr == "" {
		http.Error(w, "Отсутствуют обязательные параметры формы", http.StatusBadRequest)
		return
	}
	rateValue, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		http.Error(w, "Неверный формат курса", http.StatusBadRequest)
		return
	}
	rates, err := h.service.GetRates(r.Context())
	if err != nil {
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	var foundRate *models.ExchangeRate
	for i, r := range rates {
		if r.BaseCurrency.Code == baseCode && r.TargetCurrency.Code == targetCode {
			foundRate = &rates[i]
			break
		}
	}
	if foundRate == nil {
		http.Error(w, "Курс для указанной пары не найден", http.StatusNotFound)
		return
	}
	foundRate.Rate = rateValue
	_, err = h.service.UpdateRate(r.Context(), foundRate)
	if err != nil {
		http.Error(w, "Не удалось обновить курс в базе данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundRate)
}
