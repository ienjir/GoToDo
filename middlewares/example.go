package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add your logic here
		c.Next()
	}
}
