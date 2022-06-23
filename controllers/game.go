package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//GetGames godoc
// @Summary      Show all games
// @Description  Route to show all games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Game
// @Failure      400  {object}  httputil.HTTPError
// @Router       /games [get]
func GetGames(c *gin.Context) {
	var games []models.Game
	database.DB.Table("GAME").Find(&games)
	c.JSON(http.StatusOK, &games)
}

//GetGenres godoc
// @Summary      Show all genres
// @Description  Route to show all genres
// @Tags         genres
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Genre
// @Failure      400  {object}  httputil.HTTPError
// @Router       /genres [get]
func GetGenres(c *gin.Context) {
	var genres []models.Genre
	database.DB.Table("GENRE").Find(&genres)
	c.JSON(http.StatusOK, &genres)
}

//GetStores godoc
// @Summary      Show all stores
// @Description  Route to show all stores
// @Tags         stores
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Store
// @Failure      400  {object}  httputil.HTTPError
// @Router       /stores [get]
func GetStores(c *gin.Context) {
	var stores []models.Store
	database.DB.Table("STORE").Find(&stores)
	c.JSON(http.StatusOK, &stores)
}

//GetHistory godoc
// @Summary      Show all history
// @Description  Route to show all history
// @Tags         history
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.History
// @Failure      400  {object}  httputil.HTTPError
// @Router       /history [get]
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
