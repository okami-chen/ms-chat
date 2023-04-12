package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware(c *gin.Context) {
	c.Next()
}
