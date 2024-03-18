package middlewares

import (
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

// ZapLoggerMiddleware creates a middleware function specific to gin.
// This function logs every HTTP request.
func ZapLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// save the start time
		startTime := time.Now()

		// Process the request
		c.Next()

		// calculate the latency
		latency := time.Since(startTime)

		// determine the log level
		level := levelByStatusCode(c.Writer.Status())

		// log the request
		logger.Logger().Log(level, "incoming request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		)
	}
}

// levelByStatusCode returns the log level based on the status code.
func levelByStatusCode(statusCode int) zapcore.Level {
	var logLevel zapcore.Level
	switch {
	case statusCode >= http.StatusInternalServerError:
		logLevel = zapcore.ErrorLevel
	case statusCode >= http.StatusBadRequest:
		logLevel = zapcore.WarnLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	return logLevel
}
