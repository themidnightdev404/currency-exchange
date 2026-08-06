package models

type ExchangeRate struct {
	ID               int64   `json:"id"`
	BaseCurrencyId   int64   `json:"base_currency_id"`
	TargetCurrencyId int64   `json:"target_currency_id"`
	Rate             float64 `json:"rate"`
}
