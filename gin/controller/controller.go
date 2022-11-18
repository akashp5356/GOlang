package controller

import (
	"context"
	"fmt"
	"gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Connection = "mongodb://localhost:27017"
const dbname = "Test"
const colname = "Movietest"

var collection *mongo.Collection

func init() {
	//initialize before everything else
	clientOption := options.Client().ApplyURI(Connection)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successful")
	collection = client.Database(dbname).Collection(colname)
	fmt.Println("Collection is ready")
}

// helpers
func insertonemovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted movie with id:", inserted.InsertedID)
}
func updateonemovie(movieId string) {
	//id:=movieId
	// if err != nil {
	// 	log.Fatal(err)
	// }
	filter := bson.M{"mid": movieId}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count is :", res.ModifiedCount)
}

func deleteonemovie(movieId string) {
	//id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"mid": movieId}
	cnt, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie with count:", cnt)
}

func deleteallmovie() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted:", res.DeletedCount)
}

func getallmovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
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



//functions
func Getallmovies(c *gin.Context) {
	allmovies := getallmovies()
	c.JSON(http.StatusOK, gin.H{"result": allmovies})
}

func Createmovie(c *gin.Context) {
	var movie model.Netflix
	err := c.BindJSON(&movie)
	if err != nil {
		log.Fatal(err)
	}
	insertonemovie(movie)
	c.JSON(http.StatusOK, gin.H{"result": movie})
}

func Markwatched(c *gin.Context) {

	params := c.Param("id")
	updateonemovie(params)
	c.JSON(http.StatusOK, gin.H{"Updated": params})
}

func Deletemovie(c *gin.Context) {
	params := c.Param("id")
	deleteonemovie(params)
	c.JSON(http.StatusOK, gin.H{"deleted": params})
	// json.NewEncoder(w).Encode(params["id"])
}

func Deleteallmovies(c *gin.Context) {
	deleteallmovie()
	c.JSON(http.StatusOK, gin.H{"result": "Deleted"})
}
