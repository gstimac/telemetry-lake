package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-study/controllers"
	"go-study/middlewares"
	"io"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())

	gh := router.Group("github")
	{
		gh.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "got it!",
			})
		})
		gh.POST("event", func(c *gin.Context) {
			jsonData, err := io.ReadAll(c.Request.Body)
			fmt.Print(string(jsonData))
			if err != nil {
				c.Err()
				fmt.Print(err)
			}
			c.JSON(200, gin.H{
				"message": string(jsonData),
			})
		})
	}

	return router
}
