package main

import (
	"fmt"
	"github.com/gdscduzceuniversity/du.git/db"
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/gdscduzceuniversity/du.git/routes"
	"os"
)

func init() {
	// initialize the database connection
	db.Setup()
}

func main() {

	// Start the server
	err := routes.StartServer()
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("error starting server: %s", err))
		os.Exit(1)
	}

	// Defer the logger sync
	defer logger.Logger().Sync()
	// Defer the database disconnect
	defer db.Disconnect()
}
