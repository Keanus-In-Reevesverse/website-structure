package models

import "github.com/shopspring/decimal"

type User struct {
	UserId int64
	Name   string
	Email  string
	Phone  int64
}

type Login struct {
	UserId   int64
	Email    string
	Password string
}

type History struct {
	GameId     int
	StoreName  string
	Price      decimal.Decimal
	ChangeDate string
}
