package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

type LoggerMiddleware struct{}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		gin.DefaultWriter.Write([]byte("[GIN] " +
			endTime.Format("2006-01-02 15:04:05") + " | " +
			clientIP + " | " +
			reqMethod + " | " +
			reqURI + " | " +
			time.Now().Sub(startTime).String() + " | " +
			"Status: " + string(statusCode) + " | " +
			"Latency: " + latencyTime.String() + "\n"))
	}
}
