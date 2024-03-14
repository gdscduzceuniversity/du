package main

import (
	"github.com/gdscduzceuniversity/du.git/db"
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gdscduzceuniversity/du.git/routes"
	"os"
)

func init() {
	// Initialize the logger
	logger.InitLogger()
	// initialize the database connection
	db.Setup()
}

func main() {

	// Start the server
	err := routes.StartServer()
	if err != nil {
		logger.Sugar.Errorf("error starting server: %w", err)
		os.Exit(1)
	}

	// Defer the logger sync
	defer logger.Sugar.Sync()
	// Defer the database disconnect
	defer db.Disconnect()
}
