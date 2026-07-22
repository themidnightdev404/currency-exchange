package repositories

import (
	"currency-exchange/internal/models"
	"database/sql"
)

type CurrencyRepository struct {
	db *sql.DB
}

func (r *CurrencyRepository) FindAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	rows, err := r.db.Query("SELECT id, code, fullname, sign FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Currency
		err := rows.Scan(&c.ID, &c.Code, &c.FullName, &c.Sign)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, c)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}
