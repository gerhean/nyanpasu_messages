package main

// https://go.dev/doc/tutorial/web-service-gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var driver_URI string = ""

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// insertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func insertOne(client *mongo.Client, ctx context.Context,
	dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)
	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// album represents data about a record album.
type boardMessage struct {
	ID   float64 `json:"count"`
	Msg  string  `json:"msg"`
	Time string  `json:"time"`
}

// albums slice to seed record album data.
var messages = []boardMessage{
	{ID: 1, Msg: "Good morning!", Time: "2022-01-30T08:01:46.356Z"},
	{ID: 2, Msg: "Nyanpasu!", Time: "2022-01-30T08:03:10.862Z"},
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}

// postAlbums adds an album from JSON received in the request body.
func postMessages(c *gin.Context) {
	var newMessage boardMessage
	// value, _ := c.GetRawData()
	// fmt.Print(value)

	// Call BindJSON to bind the received JSON to
	// newMessage.
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	// Add the new album to the slice.
	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(driver_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// router := gin.Default()
	// router.GET("/messages", getMessages)
	// router.POST("/messages", postMessages)

	// router.Run("localhost:8080")
}
