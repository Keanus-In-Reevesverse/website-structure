package routes

import (
	"github.com/Keanus-In-Reevesverse/website-structure/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	//User
	r.GET("/users/:user_id", controllers.GetUserById)

	r.POST("/users", controllers.NewUser)

	//Game

	//Price

	//Run
	r.Run()
}
