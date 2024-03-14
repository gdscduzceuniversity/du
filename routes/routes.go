package routes

import (
	"fmt"
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	router.Use(ginLogger(logger.Sugar))
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
		})
		// Auth Routes
		//v1.POST("/register", handlers.Register)
		//v1.POST("/login", handlers.Login)
		//v1.POST("/logout", handlers.Logout)
		//v1.GET("/user", handlers.User)
	}

	return router
}

// ginLogger returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
func ginLogger(sugarLogger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time of the request
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Processing end time and elapsed time
		latency := time.Since(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Generate the log message format
		logMessage := fmt.Sprintf("%s %s - \"%s %s%s %d %s \"%s\" %s",
			clientIP,
			method,
			path,
			query,
			statusCode,
			latency,
			c.Request.UserAgent(),
			comment,
		)

		// Log the request with the log level based on the status code
		logByStatusCode(statusCode, sugarLogger, logMessage)
	}
}

// Set log level according to HTTP status code
func logByStatusCode(statusCode int, sugarLogger *zap.SugaredLogger, logMessage string) {
	// Set log level according to HTTP status code
	switch {
	case statusCode >= 400 && statusCode <= 599:
		// Log at Error level for 4xx errors
		sugarLogger.Errorf(logMessage)
	default:
		// Log at Info level for all other cases
		sugarLogger.Infof(logMessage)
	}
}
