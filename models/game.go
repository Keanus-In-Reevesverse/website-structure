package models

type Game struct {
	GameId      uint
	GenreId     uint
	Name        string
	Publisher   string
	Description string
	Image       string
	Video       string
}

type Genre struct {
	GenreId     uint
	Description string
}

type Store struct {
	StoreId uint
	Name    string
}
