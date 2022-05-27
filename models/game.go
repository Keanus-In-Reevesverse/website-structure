package models

import (
	"github.com/shopspring/decimal"
	"gopkg.in/validator.v2"
)

type Genre struct {
	GenreId     int64  `json:"genre_id" validate:"max=60, regexp=^[0-9]*$"`
	Description string `json:"description" validate:"nonzero, max=60"`
}

func GenreValidator(genre *Genre) error {
	if err := validator.Validate(genre); err != nil {
		return err
	}

	return nil
}

type Game struct {
	GameId      int64  `json:"game_id" validate:"regexp=^[0-9]*$"`
	GenreId     int64  `json:"genre_id" validate:"regexp=^[0-9]*$"`
	Name        string `json:"name" validate:"max=60, nonzero"`
	Publisher   string `json:"publisher" validate:"max=60, nonzero"`
	Description string `json:"description" validate:"max=300, nonzero"`
	Image       string `json:"image" validate:"max=200, nonzero"`
	Video       string `json:"video" validate:"max=60, nonzero"`
}

func GameValidator(game *Game) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}

type Store struct {
	StoreId int64  `json:"store_id" validate:"regexp=^[0-9]*$"`
	Name    string `json:"name" validate:"max=60, nonzero"`
}

func StoreValidator(store *Store) error {
	if err := validator.Validate(store); err != nil {
		return err
	}

	return nil
}

type History struct {
	GameId     int64  `json:"game_id" validate:"regexp=^[0-9]*$"`
	StoreName  string `json:"store_name" validate:"max=60, nonzero"`
	Price      decimal.Decimal
	ChangeDate string `json:"change_date" validate:"max=60, nonzero"`
}

func HistoryValidator(history *History) error {
	if err := validator.Validate(history); err != nil {
		return err
	}

	return nil
}
