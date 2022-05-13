package models

import "github.com/shopspring/decimal"

type PriceAlert struct {
	UserId int64
	GameId int64
	Price  decimal.Decimal
}

type GamePrice struct {
	GameId       int64
	StoreId      int64
	CurrentPrice decimal.Decimal
}
