package models

import (
	"github.com/shopspring/decimal"
	"gopkg.in/validator.v2"
)

type Alert struct {
	UserId int64           `json:"user_id" validate:"regexp=^[0-9]*$"`
	GameId int64           `json:"game_id" validate:"regexp=^[0-9]*$"`
	Price  decimal.Decimal `json:"price"`
}

func AlertValidator(alert *Alert) error {
	if err := validator.Validate(alert); err != nil {
		return err
	}

	return nil
}
