package models

type Game struct {
	GameId      int64
	GenreId     int64
	Name        string
	Publisher   string
	Description string
	Image       string
	Video       string
}

type Genre struct {
	GenreId     int64
	Description string
}

type Store struct {
	StoreId int64
	Name    string
}
