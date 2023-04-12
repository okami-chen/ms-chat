package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(c *gin.Context, fillInterval time.Duration, cap, quantum int64) {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	if bucket.TakeAvailable(1) < 1 {
		c.String(http.StatusForbidden, "rate limit...")
		c.Abort()
		return
	}
	c.Next()
}
