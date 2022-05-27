package tests

import (
	"github.com/Keanus-In-Reevesverse/website-structure/database"
	"github.com/Keanus-In-Reevesverse/website-structure/models"
	"github.com/gin-gonic/gin"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	user := models.User{UserId: 102, Name: "User Name Test",
		Email: "test@mail.com", PhoneNumber: "12966667777", Password: "senha1234"}

	database.DB.Create(&user)
	ID = int(user.UserId)
}

func DeleteStudentMock() {
	var user models.User
	database.DB.Delete(&user, ID)
}
