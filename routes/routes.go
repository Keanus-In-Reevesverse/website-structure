package routes

import (
	"github.com/Keanus-In-Reevesverse/website-structure/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	//User
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:user_id", controllers.GetUser)
	r.GET("/users/login", controllers.GetLogin)
	r.GET("/user/login/:user_id", controllers.GetLogins)

	r.POST("/users", controllers.CreateUser)

	//Game

	//Price

	//Run
	r.Run()
}
