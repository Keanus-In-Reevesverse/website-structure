package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Create
func CreateAlert(c *gin.Context) {
	var alert models.Alert

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na criação": err.Error()})
		return
	}

	if err := models.AlertValidator(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na request": err.Error()})
		return
	}

	alert = database.AlertOps(alert, "create")

	c.JSON(http.StatusOK, &alert)
}

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

//GetAlerts godoc
// @Summary      Show all alerts
// @Description  Route to show all alerts
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Alert
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alerts [get]
func GetAlerts(c *gin.Context) {
	var alerts []models.Alert
	database.DB.Table("PRICE_ALERT").Find(&alerts)
	c.JSON(http.StatusOK, &alerts)
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

func GetAlertsByIds(c *gin.Context) {
	var alerts []models.Alert
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
