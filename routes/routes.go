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
	r.GET("/games", controllers.GetGames)
	r.GET("/games/:game_id", controllers.GetGameById)

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
	r.GET("/prices/:store_id", controllers.GetPricesByStoreId)

	//Alert
	r.GET("/alerts", controllers.GetPriceAlerts)
	r.GET("/alerts/:user_id/:game_id", controllers.GetAlertsByIds)

	//Run
	r.Run()
}
