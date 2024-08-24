package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"telemetry-lake/internal/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		var secret = cfg.GetString("github.webhook.secret")
		var sigHeader = cfg.GetString("github.webhook.signature-header")
		hmacHeader := c.Request.Header.Get(sigHeader)

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		// Write Body back to request if we need to use it later
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		splitHeader := strings.Split(hmacHeader, "=")
		hashAlg := splitHeader[0]
		messageMac := splitHeader[1]

		if hashAlg != "sha256" {
			c.AbortWithStatus(500)
		}
		if len(messageMac) == 0 {
			c.AbortWithStatus(401)
		}

		messageMacBuf, _ := hex.DecodeString(messageMac)
		if !ValidMAC(bodyBytes, messageMacBuf, []byte(secret)) {
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
	return hmac.Equal(expectedMAC, messageMAC)
}
