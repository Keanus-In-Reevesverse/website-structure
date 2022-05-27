package models

import (
	"github.com/shopspring/decimal"
	"gopkg.in/validator.v2"
)

type PriceAlert struct {
	UserId int64 `json:"user_id" validate:"regexp=^[0-9]*$"`
	GameId int64 `json:"game_id" validate:"regexp=^[0-9]*$"`
	Price  decimal.Decimal
}

func PriceAlertValidator(priceAlert *PriceAlert) error {
	if err := validator.Validate(priceAlert); err != nil {
		return err
	}

	return nil
}

type GamePrice struct {
	GameId       int64  `json:"game_id" validate:"regexp=^[0-9]*$"`
	StoreName    string `json:"store_name" validate:"max=60, nonzero"`
	CurrentPrice decimal.Decimal
}

func GamePriceValidator(gamePrice *GamePrice) error {
	if err := validator.Validate(gamePrice); err != nil {
		return err
	}

	return nil
}
