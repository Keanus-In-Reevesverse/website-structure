package models

import "github.com/shopspring/decimal"

type User struct {
	UserId uint
	Name   string
	Email  string
	Phone  int64
}

type Login struct {
	UserId   uint
	Email    string
	Password string
}

type History struct {
	GameId     uint
	StoreName  string
	Price      decimal.Decimal
	ChangeDate string
}

type Favorite struct {
	UserId uint
	GameId uint
}
