package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/okamin-chen/service/pkg/global"
	"time"
)

var skip map[string]struct{}

func LoggerMiddleware(c *gin.Context) {
	l := c.Copy()

	start := time.Now()
	path := l.Request.URL.Path
	raw := l.Request.URL.RawQuery

	param := gin.LogFormatterParams{
		Request:   l.Request,
		Keys:      l.Keys,
		TimeStamp: time.Now(),
	}
	// Process request
	c.Next()

	// Log only when path is not being skipped
	if _, ok := skip[path]; !ok {

		// Stop timer
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = l.ClientIP()
		param.Method = l.Request.Method
		param.StatusCode = l.Writer.Status()
		param.ErrorMessage = l.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = l.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		global.Log.Info(fmt.Sprintf("%s %s %s %d \"%s\" %s \"%s\" %s\n",
			param.ClientIP,
			param.Request.Proto,
			param.Method,
			param.StatusCode,
			param.Path,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		))

	}
}
