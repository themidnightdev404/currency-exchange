package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/themidnightdev404/currency-exchange/internal/models"
	"github.com/themidnightdev404/currency-exchange/internal/repositories"
)

var ErrCurrencyNotFound = errors.New("currency not found")

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
	currency, err := r.repo.GetCurrencyByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCurrencyNotFound
		}
		return nil, err
	}
	return currency, nil
}

func (r *CurrencyService) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	return r.repo.GetCurrencies(ctx)
}
