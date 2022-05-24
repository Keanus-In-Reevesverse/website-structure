package models

import "github.com/shopspring/decimal"

type User struct {
	UserId      int64
	Name        string
	Email       string
	PhoneNumber string
}

type UserInput struct {
	Name        string
	Email       string
	Password    string
	PhoneNumber string
}

type UserOutput struct {
	Name        string
	Email       string
	PhoneNumber string
}
type History struct {
	GameId     int64
	StoreName  string
	Price      decimal.Decimal
	ChangeDate string
}

type Favorite struct {
	UserId int64
	GameId int64
}
