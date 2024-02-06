package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"telemetry-lake/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		var secret = cfg.GetString("github.webhook.secret")
		var sigHeader = cfg.GetString("github.webhook.signature-header")

		hmacHeader := c.Request.Header.Get(sigHeader)
		byteBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Err()
		}

		splitHeader := strings.Split(hmacHeader, "=")
		algorithm := splitHeader[0]
		signature := splitHeader[1]

		if algorithm != "sha256" {
			c.AbortWithStatus(500)
		}
		if len(signature) == 0 {
			c.AbortWithStatus(401)
		}
		if !ValidMAC(string(byteBody), signature, secret) {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func ValidMAC(message string, messageMAC string, key string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return messageMAC == expectedMAC
}
