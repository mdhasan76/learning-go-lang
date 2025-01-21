package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// const connectionString = "mongodb+srv://shohag_isp:oA1zfB43JZkFLHqx@cluster0.di4ojvf.mongodb.net/"
const connectionString = "mongodb+srv://go_netflix:lW28H9FzpWhDA2gJ@cluster0.mj4ed9j.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const collectionName = "watchlist"

// most important
var collection *mongo.collection

// Connect with db
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}
