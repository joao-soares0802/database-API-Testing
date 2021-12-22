package routes

import (
	"awesomeProject/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1/")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("", controllers.CreateBook)
			books.PUT("/update", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		games := main.Group("games")
		{
			games.GET("/:id", controllers.ShowGame)
			games.GET("/", controllers.ShowGames)
			games.POST("", controllers.CreateGame)
			games.PUT("/update", controllers.UpdateGames)
			games.DELETE("/:id", controllers.DeleteGame)

		}
	}
	return router
}
