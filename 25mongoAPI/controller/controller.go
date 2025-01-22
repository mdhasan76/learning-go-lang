package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoAPI/model"
	"net/http"

	"github.com/gorilla/mux"
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

func insertOneMovie(movie model.Netflix) primitive.M {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	var newMovie primitive.M
	err = collection.FindOne(context.Background(), bson.M{"_id": inserted.InsertedID}).Decode(&newMovie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
	return newMovie
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

func deleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all Movie successfully. Total deleted count is: ", result.DeletedCount, nil)
	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func getOneMovie(movieId string) (primitive.M, string) {
	var movie primitive.M
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	res := collection.FindOne(context.Background(), filter)
	err := res.Decode(&movie)
	if err != nil {
		return movie, "Data not found"
	}
	return movie, ""
}

// Actual controller -file

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movieId := mux.Vars(r)["id"]
	data, msg := getOneMovie(movieId)
	if msg != "" {
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
		return
	}
	json.NewEncoder(w).Encode(data)
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	newMovie := insertOneMovie(movie)
	json.NewEncoder(w).Encode(newMovie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	param := mux.Vars(r)["id"]
	updateOneMovie(param)
	json.NewEncoder(w).Encode(param)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	movieId := mux.Vars(r)["id"]
	deleteOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	w.Write([]byte(fmt.Sprintf("All movies Deleted successfully. and deleted count is: %v", count)))
}
