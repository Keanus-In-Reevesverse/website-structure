package models

type Game struct {
	GameId      int
	GenreId     int
	Name        string
	Publisher   string
	Description string
	Image       string
	Video       string
}

type Favorite struct {
	UserId int
	GameId int
}

type Genre struct {
	GenreId     int
	Description string
}

type Store struct {
	StoreId int
	Name    string
}
