package controller

import (
	"context"
	"fmt"
	"log"
	model "mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// define db url whether local or remote url, dbname & collection name
const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collectionName = "watchlist"

/*
* most important
* declare a var. which will be storing reference to mongodb collection
 */
var collection *mongo.Collection

/*
* create a init method
* which is specialized method in golang
* which runs only at the 1 time when this application is going to execute
* and will run only 1 time
 */

func init() {
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
	//else connection success
	fmt.Println("mongodb connection successfull")

	//now create db & collection & store it's reference in the var.
	//collection which we decalred above
	collection = client.Database(dbName).Collection(collectionName)

	//if collection reference or instance is ready just print it
	fmt.Println("collection instance is ready")
}

//Mongodb helpers function

//insert one record

func insertOne(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one record with id: ", inserted.InsertedID)
}

// update movie
func updateMovie(movieId string) {
	//convert movieId string into mongodb id
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	//create filter from above id
	filter := bson.M{"_id": id}
	deletedCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deletedCount)
}

func deleteAllMovies() int64 {
	filter := bson.D{{}}
	deletedResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count are: ", deletedResult.DeletedCount)
	return deletedResult.DeletedCount
}

// get all movie
func getAllMovie() []primitive.D {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.D
	for cur.Next(context.Background()) {
		var movie bson.D
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}
