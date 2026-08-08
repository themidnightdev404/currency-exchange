package service

import (
	"context"
	"errors"

	"github.com/themidnightdev404/currency-exchange/internal/models"
	"github.com/themidnightdev404/currency-exchange/internal/repositories"
)

type ExchangeRateService struct {
	repo         *repositories.ExchangeRateRepository
	currencyRepo *repositories.CurrencyRepository
}

func NewExchangeRateService(repo *repositories.ExchangeRateRepository, currencyRepo *repositories.CurrencyRepository) *ExchangeRateService {
	return &ExchangeRateService{
		repo:         repo,
		currencyRepo: currencyRepo,
	}
}

func (r *ExchangeRateService) CreateRate(ctx context.Context, rate *models.ExchangeRate) (*models.ExchangeRate, error) {
	if rate.Rate <= 0 {
		return nil, errors.New("Курс обмена должен быть больше нуля")
	}
	return r.repo.CreateRate(ctx, rate)
}

func (r *ExchangeRateService) UpdateRate(ctx context.Context, rate *models.ExchangeRate) (*models.ExchangeRate, error) {
	if rate.Rate <= 0 {
		return nil, errors.New("Курс обмена должен быть больше нуля")
	}
	return r.repo.UpdateRate(ctx, rate)
}

func (r *ExchangeRateService) GetRate(ctx context.Context, id int64) (*models.ExchangeRate, error) {
	return r.repo.GetRate(ctx, id)
}

func (r *ExchangeRateService) GetRates(ctx context.Context) ([]models.ExchangeRate, error) {
	return r.repo.GetRates(ctx)
}
