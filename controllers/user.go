package controllers

import (
	"net/http"

	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.First(&users)
	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("userid")

	database.DB.First(&user, id)

	if user.UserId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
