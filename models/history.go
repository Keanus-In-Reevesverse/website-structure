package models

import "github.com/shopspring/decimal"

type History struct {
	GameId     int64  `json:"game_id"`
	StoreName  string `json:"store_name"`
	Price      decimal.Decimal
	ChangeDate string `json:"change_date"`
}
