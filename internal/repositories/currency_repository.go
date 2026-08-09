package repositories

import (
	"context"
	"database/sql"

	"github.com/themidnightdev404/currency-exchange/internal/models"
)

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) CreateCurrency(ctx context.Context, currency *models.Currency) (*models.Currency, error) {
	query := `
		INSERT INTO currencies (Code, Name, Sign)
		VALUES (?, ?, ?)
		RETURNING ID;
	`
	var lastInsertID int64
	err := r.db.QueryRowContext(ctx, query, currency.Code, currency.FullName, currency.Sign).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	currency.ID = lastInsertID
	return currency, nil
}

func (r *CurrencyRepository) GetCurrencyByCode(ctx context.Context, code string) (*models.Currency, error) {
	query := `
		SELECT ID, Code, Name, Sign FROM currencies WHERE Code = ?;
	`
	var currency models.Currency
	err := r.db.QueryRowContext(ctx, query, code).Scan(&currency.ID, &currency.Code, &currency.FullName, &currency.Sign)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *CurrencyRepository) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	query := `
		SELECT ID, Code, Name, Sign FROM currencies;
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var currencies []models.Currency
	for rows.Next() {
		var c models.Currency
		err := rows.Scan(&c.ID, &c.Code, &c.FullName, &c.Sign)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return currencies, nil
}
