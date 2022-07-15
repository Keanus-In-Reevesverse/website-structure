package routes

import (
	"github.com/Keanus-In-Reevesverse/website-structure/controllers"
	"github.com/Keanus-In-Reevesverse/website-structure/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	//Swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Home Access
	r.GET("/home", controllers.Home)

	//Register
	r.POST("/users", controllers.NewUser)
	r.PATCH("/users", controllers.EditUser)
	r.DELETE("/users", controllers.DeleteUser)

	//Login
	r.POST("/login", controllers.Login)

	//Game
	r.GET("/games", controllers.GetGames)
	r.GET("/games/:game_id", controllers.GetGameById)

	//Favorite
	r.GET("/favorites", controllers.GetFavorites)
	r.GET("/favorites/:user_id", controllers.GetFavoritesByUserId)

	//Genre
	r.GET("/genres", controllers.GetGenres)
	r.GET("/genres/:genre_id", controllers.GetGenreById)

	//Stores
	r.GET("/stores", controllers.GetStores)
	r.GET("/stores/:store_id", controllers.GetStoreById)

	//History
	r.GET("/history", controllers.GetHistory)
	r.GET("/history/:game_id", controllers.GetHistoryByGameId)

	//Price
	r.GET("/prices", controllers.GetPrices)
	r.GET("/prices/:game_id", controllers.GetPricesByGameId)

	//Alert
	r.POST("/alerts", controllers.CreateAlert)
	r.PATCH("/alerts", controllers.EditAlert)
	r.DELETE("/alerts", controllers.DeleteAlert)
	r.GET("/alerts", controllers.GetAlerts)
	r.GET("/alerts/:user_id/:game_id", controllers.GetAlertsByIds)

	//Run
	r.Run()
}
