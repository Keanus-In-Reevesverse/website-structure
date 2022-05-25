package routes

import (
	"github.com/Keanus-In-Reevesverse/website-structure/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	//Register
	r.GET("/users/:user_id", controllers.GetUserById)
	r.POST("/users", controllers.NewUser)

	//Login
	r.POST("/login", controllers.Login)

	//Game

	//Price

	//Run
	r.Run()
}
