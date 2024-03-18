package db

import (
	"context"
	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	databseName         = "du"
	Client              *mongo.Client
	UserCollection      *mongo.Collection
	CommunityCollection *mongo.Collection
	PostCollection      *mongo.Collection
)

func Setup() {
	// load .env file to get the connection string for the database
	// it is optional to use .env file, you can use the environment variable rather than using .env file
	// if you want to use .env file, you can create a .env file to root of the project and add the connection string
	// etc. MONGODB_URI=mongodb://localhost:27017
	if err := godotenv.Load("du.env"); err != nil {
		logger.Sugar.Info("No .env file found")
	}
	// if you want to use the environment variable, you can set the environment variable
	// in your terminal
	// etc. export MONGODB_URI=mongodb://localhost:27017
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		logger.Sugar.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	// Create a new Client and connect to the server
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	logger.Sugar.Info("Pinged your deployment. You successfully connected to MongoDB!")

	getCollections()
}

// getCollections from the database and assign them to the variables
func getCollections() {
	database := Client.Database(databseName)

	UserCollection = database.Collection("user")
	CommunityCollection = database.Collection("community")
	PostCollection = database.Collection("post")
}

// Disconnect from the server when the application exits
func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	logger.Sugar.Info("Connection to MongoDB closed.")
}
