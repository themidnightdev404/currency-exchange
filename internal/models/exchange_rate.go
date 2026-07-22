package models

import (
	"github.com/shopspring/decimal"
)

type ExchangeRate struct {
	ID               int
	BaseCurrencyID   int
	TargetCurrencyID int
	Rate             decimal.Decimal
}
