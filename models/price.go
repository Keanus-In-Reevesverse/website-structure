package models

import "github.com/shopspring/decimal"

type PriceAlert struct {
	UserId uint
	GameId uint
	Price  decimal.Decimal
}

type GamePrice struct {
	GameId       uint
	StoreId      uint
	CurrentPrice decimal.Decimal
}
