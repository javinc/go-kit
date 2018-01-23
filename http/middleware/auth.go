package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth middleware using JWT
func Auth(callback func(*gin.Context, string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		headerStr := strings.TrimSpace(c.GetHeader("authorization"))
		if headerStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "bearer token required",
			})

			return
		}

		// check for bearer
		authHeaders := strings.Split(headerStr, " ")
		if len(authHeaders) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization token",
			})
		}

		// validate token
		err := callback(c, authHeaders[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		// continue
		c.Next()
	}
}
