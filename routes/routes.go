package routes

import (
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gdscduzceuniversity/du.git/middlewares"
	"github.com/gin-gonic/gin"
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
	router.Use(middlewares.ZapLoggerMiddleware())
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
