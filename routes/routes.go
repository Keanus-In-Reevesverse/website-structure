package routes

import (
	"github.com/Keanus-In-Reevesverse/website-structure/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	//User
	r.GET("/users", controllers.GetUsers)
	r.GET("/user:id", controllers.GetUser)
	//Game

	//Price
}
