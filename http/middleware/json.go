package middleware

import (
	"github.com/gin-gonic/gin"
)

// JSON middleware to force Content-Type to JSON
func JSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
	}
}
