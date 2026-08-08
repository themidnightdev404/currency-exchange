package service

import (
	"context"
	"errors"

	"github.com/themidnightdev404/currency-exchange/internal/models"
	"github.com/themidnightdev404/currency-exchange/internal/repositories"
)

type CurrencyService struct {
	repo *repositories.CurrencyRepository
}

func NewCurrencyService(repo *repositories.CurrencyRepository) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (r *CurrencyService) CreateCurrency(ctx context.Context, currency *models.Currency) (*models.Currency, error) {
	if len(currency.Code) != 3 {
		return nil, errors.New("Код валюты должен содержать ровно 3 символа")
	}
	if currency.FullName == "" || currency.Sign == "" {
		return nil, errors.New("Поле не может быть пустым")
	}
	return r.repo.CreateCurrency(ctx, currency)
}

func (r *CurrencyService) GetCurrencyByCode(ctx context.Context, code string) (*models.Currency, error) {
	return r.repo.GetCurrencyByCode(ctx, code)
}

func (r *CurrencyService) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	return r.repo.GetCurrencies(ctx)
}
