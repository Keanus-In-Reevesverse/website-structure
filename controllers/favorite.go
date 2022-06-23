package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

//GetFavorites godoc
// @Summary      Show all favorites
// @Description  Route to show all favorites
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Favorite
// @Failure      400  {object}  httputil.HTTPError
// @Router       /favorites [get]
func GetFavorites(c *gin.Context) {
	var favorites []models.Favorite
	database.DB.Table("PRICE_ALERT").Find(&favorites)
	c.JSON(http.StatusOK, &favorites)
}

//GetFavoriteByUserId godoc
// @Summary      Show favorite by user_id
// @Description  Route to show favorite game by user_id
// @Tags         favorites
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Favorite
// @Failure      404  {object}  httputil.HTTPError
// @Router       /favorites/id [get]
func GetFavoritesByUserId(c *gin.Context) {
	var favorites []models.Favorite
	id := c.Params.ByName("user_id")

	database.DB.Table("FAVORITES").Find(&favorites).Where("user_id = ?", id)

	for _, favorite := range favorites {
		if favorite.GameId == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Not Found": "Game not found"})
			return
		}
	}

	c.JSON(http.StatusOK, favorites)
}
