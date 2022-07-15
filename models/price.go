package models

import (
	"github.com/shopspring/decimal"
)

type Price struct {
	GameId       int64           `json:"game_id"`
	StoreName    string          `json:"store_name"`
	CurrentPrice decimal.Decimal `json:"current_price"`
}
