package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	var favorites []models.Favorite
	database.DB.Table("PRICE_ALERT").Find(&favorites)
	c.JSON(http.StatusOK, &favorites)
}

func GetFavoritesByUserId(c *gin.Context) {
	var favorites []models.GamePrice
	id := c.Params.ByName("user_id")

	database.DB.Table("FAVORITES").Find(&favorites).Where("user_id = ?", id)

	for _, favorite := range favorites {
		if favorite.GameId == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "Game not found"})
			return
		}
		if favorite.StoreName == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "Store not found"})
			return
		}
	}

	c.JSON(http.StatusOK, favorites)
}
