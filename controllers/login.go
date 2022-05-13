package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

//Getters
func GetLogins(c *gin.Context) {
	var login []models.Login
	database.DB.Table("LOGIN").Find(&login)
	c.JSON(200, &login)
}

func GetLoginById(c *gin.Context) {
	var login models.Login
	id := c.Params.ByName("user_id")

	database.DB.Table("LOGIN").First(&login, id)

	if login.UserId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, login)
}
