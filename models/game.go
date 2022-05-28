package models

import (
	"github.com/shopspring/decimal"
)

type Genre struct {
	GenreId     int64  `json:"genre_id"`
	Description string `json:"description"`
}

type Game struct {
	GameId      int64  `json:"game_id"`
	GenreId     int64  `json:"genre_id"`
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

type Store struct {
	StoreId int64  `json:"store_id"`
	Name    string `json:"name"`
}

type History struct {
	GameId     int64  `json:"game_id"`
	StoreName  string `json:"store_name"`
	Price      decimal.Decimal
	ChangeDate string `json:"change_date"`
}
