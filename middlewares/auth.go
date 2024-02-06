package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//config := config.GetConfig()
		hmacHeader := c.Request.Header.Get("x-hub-signature-256")
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Err()
		}

		message := string(jsonData)
		splitHeader := strings.Split(hmacHeader, "=")
		algorithm := splitHeader[0]
		signature := splitHeader[1]

		var secret string = "verysecret"

		if algorithm != "sha256" {
			c.AbortWithStatus(500)
		}
		if len(signature) == 0 {
			c.AbortWithStatus(401)
		}
		if !ValidMAC([]byte(message), []byte(hmacHeader), []byte(secret)) {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
