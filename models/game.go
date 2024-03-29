package models

type Game struct {
	GameId      int64  `json:"game_id"`
	GenreId     int64  `json:"genre_id"`
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}
