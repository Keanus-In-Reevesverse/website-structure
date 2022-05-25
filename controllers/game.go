package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Get arrays
func GetGames(c *gin.Context) {
	var games []models.Game
	database.DB.Table("GAME").Find(&games)
	c.JSON(http.StatusOK, &games)
}

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	database.DB.Table("GENRE").Find(&genres)
	c.JSON(http.StatusOK, &genres)
}

func GetStores(c *gin.Context) {
	var stores []models.Store
	database.DB.Table("STORE").Find(&stores)
	c.JSON(http.StatusOK, &stores)
}

func GetHistory(c *gin.Context) {
	var history []models.History
	database.DB.Table("HISTORY").Find(&history)
	c.JSON(http.StatusOK, &history)
}

//Get by ID
func GetGameById(c *gin.Context) {
	var game models.Game
	id := c.Params.ByName("game_id")

	database.DB.Table("GAME").First(&game, id)

	if game.GameId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func GetGenreById(c *gin.Context) {
	var genre models.Genre
	id := c.Params.ByName("genre_id")

	database.DB.Table("GENRE").First(&genre, id)

	if genre.GenreId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func GetStoreById(c *gin.Context) {
	var store models.Store
	id := c.Params.ByName("store_id")

	database.DB.Table("STORE").First(&store, id)

	if store.StoreId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func GetHistoryByGameId(c *gin.Context) {
	var history []models.History
	id := c.Params.ByName("game_id")

	database.DB.Table("HISTORY").Find(&history).Where("game_id = ?", id)

	for i, hist := range history {
		if hist.GameId == 0 {
			history = append(history[:i], history[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, history)
}
