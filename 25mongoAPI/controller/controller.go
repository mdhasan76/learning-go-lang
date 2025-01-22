package controller

import (
	"context"
	"fmt"
	"log"
	"mongoAPI/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

// const connectionString = "mongodb+srv://shohag_isp:oA1zfB43JZkFLHqx@cluster0.di4ojvf.mongodb.net/"
const connectionString = "mongodb+srv://go_netflix:lW28H9FzpWhDA2gJ@cluster0.mj4ed9j.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const collectionName = "watchlist"

// most important
var collection *mongo.Collection

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

// Mongodb helper  - file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie is deleted successfully", result)
}

func deleteAllMovie() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all Movie successfully. Total deleted count is: ", result.DeletedCount, nil)
	return result.DeletedCount
}
