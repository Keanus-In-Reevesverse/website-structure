package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//GetPrices godoc
// @Summary      Show all prices
// @Description  Route to show all prices
// @Tags         prices
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Price
// @Failure      400  {object}  httputil.HTTPError
// @Router       /prices [get]
func GetPrices(c *gin.Context) {
	var prices []models.Price
	database.DB.Table("GAME_PRICES").Find(&prices)
	c.JSON(http.StatusOK, &prices)
}

//Get by id
func GetPricesByGameId(c *gin.Context) {
	var prices []models.Price
	id := c.Params.ByName("game_id")

	database.DB.Table("GAME_PRICES").Find(&prices).Where("game_id = ?", id)

	for _, price := range prices {
		if price.GameId == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "Game not found"})
			return
		}
		if price.StoreName == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "Store not found"})
			return
		}
	}

	c.JSON(http.StatusOK, prices)
}
