package router

import (
	"api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://qrcode-reader-vert.vercel.app", "https://koyofes2023-reception.vercel.app"}
	router.Use(cors.New(config))

	router.GET("/users/:uid", controllers.GetReceptionUserInfo)
	router.PUT("/users/:uid", controllers.SetReceptionUserBodyTemperature)

	router.GET("/env-time", controllers.GetEnvironmentTime)
	return router
}
