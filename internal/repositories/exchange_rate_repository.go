package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/themidnightdev404/currency-exchange/internal/models"
)

type ExchangeRateRepository struct {
	db *sql.DB
}

func NewExchangeRateRepository(db *sql.DB) *ExchangeRateRepository {
	return &ExchangeRateRepository{db: db}
}

func (r *ExchangeRateRepository) CreateRate(ctx context.Context, rate *models.ExchangeRate) (*models.ExchangeRate, error) {
	query := `
		INSERT INTO ExchangeRates (BaseCurrencyId, TargetCurrencyId, Rate) 
		VALUES (?, ?, ?)
		RETURNING ID;
	`

	var lastInsertID int64
	err := r.db.QueryRowContext(ctx, query, rate.BaseCurrency.ID, rate.TargetCurrency.ID, rate.Rate).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}

	rate.ID = lastInsertID
	return rate, nil
}

func (r *ExchangeRateRepository) GetRate(ctx context.Context, id int64) (*models.ExchangeRate, error) {
	query := `
		SELECT 
			er.ID, er.Rate,
			base.ID, base.Code, base.FullName, base.Sign,
			target.ID, target.Code, target.FullName, target.Sign
		FROM ExchangeRates er
		JOIN currencies base ON er.BaseCurrencyId = base.ID
		JOIN currencies target ON er.TargetCurrencyId = target.ID
		WHERE er.ID = ?;
	`

	var rate models.ExchangeRate

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rate.ID, &rate.Rate,
		&rate.BaseCurrency.ID, &rate.BaseCurrency.Code, &rate.BaseCurrency.FullName, &rate.BaseCurrency.Sign,
		&rate.TargetCurrency.ID, &rate.TargetCurrency.Code, &rate.TargetCurrency.FullName, &rate.TargetCurrency.Sign)

	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func (r *ExchangeRateRepository) GetRates(ctx context.Context) ([]models.ExchangeRate, error) {
	query := `
		SELECT 
			er.ID, er.Rate,
			base.ID, base.Code, base.FullName, base.Sign,
			target.ID, target.Code, target.FullName, target.Sign
		FROM ExchangeRates er
		JOIN currencies base ON er.BaseCurrencyId = base.ID
		JOIN currencies target ON er.TargetCurrencyId = target.ID;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rates []models.ExchangeRate
	for rows.Next() {
		var er models.ExchangeRate
		err := rows.Scan(&er.ID, &er.Rate,
			&er.BaseCurrency.ID, &er.BaseCurrency.Code, &er.BaseCurrency.FullName, &er.BaseCurrency.Sign,
			&er.TargetCurrency.ID, &er.TargetCurrency.Code, &er.TargetCurrency.FullName, &er.TargetCurrency.Sign)
		if err != nil {
			return nil, err
		}
		rates = append(rates, er)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rates, nil
}

func (r *ExchangeRateRepository) UpdateRate(ctx context.Context, rate *models.ExchangeRate) (*models.ExchangeRate, error) {
	query := `
		UPDATE ExchangeRates
		SET BaseCurrencyId = ?, TargetCurrencyId = ?, Rate = ?
		WHERE ID = ?;
	`
	result, err := r.db.ExecContext(ctx, query, rate.BaseCurrency.ID, rate.TargetCurrency.ID, rate.Rate, rate.ID)
	if err != nil {
		return nil, err
	}

	rowsCount, _ := result.RowsAffected()
	if rowsCount == 0 {
		return nil, fmt.Errorf("курс обмена с id %d не найден", rate.ID)
	}
	return rate, nil
}
