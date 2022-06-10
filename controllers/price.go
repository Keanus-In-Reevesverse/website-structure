package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Create
func CreatePriceAlert(c *gin.Context) {
	var alert models.PriceAlert

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	if err := models.PriceAlertValidator(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	alert = database.AlertOps(alert, "create")

	c.JSON(http.StatusOK, &alert)
}

//Get arrays
func GetPrices(c *gin.Context) {
	var prices []models.GamePrice
	database.DB.Table("GAME_PRICES").Find(&prices)
	c.JSON(http.StatusOK, &prices)
}

func GetPriceAlerts(c *gin.Context) {
	var alerts []models.PriceAlert
	database.DB.Table("PRICE_ALERT").Find(&alerts)
	c.JSON(http.StatusOK, &alerts)
}

//Get by id
func GetPricesByGameId(c *gin.Context) {
	var prices []models.GamePrice
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

func GetAlertsByIds(c *gin.Context) {
	var alerts []models.PriceAlert
	user_id := c.Params.ByName("user_id")
	game_id := c.Params.ByName("game_id")

	if game_id != "" {
		database.DB.Table("ALERT_PRICES").Find(&alerts).Where("user_id = ?", user_id)
	} else {
		database.DB.Table("ALERT_PRICES").Find(&alerts).Where("user_id = ? AND game_id = ?", user_id, game_id)
	}

	for i, alert := range alerts {
		if alert.UserId == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "User Not Found"})
			return
		}
		if alert.GameId == 0 {
			alerts = append(alerts[:i], alerts[i+1:]...)
		}
	}

	c.JSON(http.StatusOK, alerts)
}

//Edit
func EditAlert(c *gin.Context) {
}

func DeleteAlert(c *gin.Context) {
}
