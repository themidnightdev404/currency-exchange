package models

type ExchangeRate struct {
	ID             int64    `json:"id"`
	BaseCurrency   Currency `json:"base_currency"`
	TargetCurrency Currency `json:"target_currency"`
	Rate           float64  `json:"rate"`
}
