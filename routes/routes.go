package routes

import (
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

const port = ":8080"

func StartServer() error {
	router := SetupRouter()
	err := router.Run(port)
	return err
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	// Custom logger middleware to log all requests to the console and the file with zap logger
	router.Use(zapLoggerMiddleware())
	router.Use(gin.Recovery())

	// for swagger api documentation group
	//docs.SwaggerInfo.BasePath = "/api/v1"
	// Swagger Route for connect to swagger ui go to http://127.0.0.1:8080/swagger/index.html
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")
	{
		// Ping Route
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			logger.Logger().Info("pong")
		})
		// Auth Routes
		//v1.POST("/register", handlers.Register)
		//v1.POST("/login", handlers.Login)
		//v1.POST("/logout", handlers.Logout)
		//v1.GET("/user", handlers.User)
	}

	return router
}

// zapLoggerMiddleware creates a middleware function specific to gin.
// This function logs every HTTP request.
func zapLoggerMiddleware() gin.HandlerFunc {
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
