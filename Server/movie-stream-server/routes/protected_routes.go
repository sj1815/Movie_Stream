package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/sj1815/MovieStream/Server/movie-stream-server/controllers"
	"github.com/sj1815/MovieStream/Server/movie-stream-server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
}
