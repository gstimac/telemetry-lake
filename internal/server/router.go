package server

import (
	"github.com/gin-gonic/gin"
	"telemetry-lake/internal/controllers"
	"telemetry-lake/internal/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthCtrl := new(controllers.HealthController)
	router.GET("/health", healthCtrl.Status)
	router.Use(middlewares.AuthMiddleware())

	ghCtrl := new(controllers.GithubController)
	gh := router.Group("github")
	{
		gh.GET("ping", ghCtrl.Ping)
		gh.POST("event", ghCtrl.WriteEvent)
	}

	return router
}
