package db

import (
	"context"
	"fmt"
	"os"

	"github.com/gdscduzceuniversity/du.git/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// databseName is the name of the database
	database = "du"
	// userCollection is the name of the collection that stores the user data
	user = "user"
	// communityCollection is the name of the collection that stores the community data
	community = "community"
	// postCollection is the name of the collection that stores the post data
	communityContent = "communityContent"
)

var (
	db                  *mongo.Database
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
		logger.Logger().Info("No .env file found")
	}
	// if you want to use the environment variable, you can set the environment variable
	// in your terminal
	// etc. export MONGODB_URI=mongodb://localhost:27017
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		logger.Logger().Fatal("You must set your 'MONGO_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
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
	logger.Logger().Info("Pinged your deployment. You successfully connected to MongoDB!")

	getDatabase()
	// Create collections
	createCollections(user, community, communityContent)
	getCollections()
}

// getDatabase from the client and assign it to the variable
func getDatabase() {
	db = Client.Database(database)
	if db == nil {
		logger.Logger().Fatal("failed to get database")
	}
}

// createCollections in the database
func createCollections(collections ...string) {
	for _, collection := range collections {
		if err := db.CreateCollection(context.TODO(), collection); err != nil {
			logger.Logger().Fatal(fmt.Sprintf("Error creating collection %s, %v:", collection, err))
		}
	}
}

// getCollections from the database and assign them to the variables
func getCollections() {
	UserCollection = db.Collection(user)
	CommunityCollection = db.Collection(community)
	PostCollection = db.Collection(community)
}

// Disconnect from the server when the application exits
func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	logger.Logger().Info("Connection to MongoDB closed.")
}
