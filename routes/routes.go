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
	r.GET("/users/login", controllers.GetLogins)
	r.GET("/user/login/:user_id", controllers.GetLogin)

	r.POST("/users", controllers.NewUser)

	//Game

	//Price

	//Run
	r.Run()
}
