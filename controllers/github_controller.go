package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type GithubController struct{}

func (h GithubController) WriteEvent(c *gin.Context) {
	byteBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Err()
		log.Print("Error reading request body")
	}
	c.JSON(200, gin.H{
		"message": string(byteBody),
	})
}

func (h GithubController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
