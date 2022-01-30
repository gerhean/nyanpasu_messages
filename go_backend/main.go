package main

// https://go.dev/doc/tutorial/web-service-gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
// context.CancelFunc will be used to cancel context and resource associated with it.
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

type boardMessage struct {
	Count float64 `json:"count,omitempty" bson:"count,omitempty"`
	Msg   string  `json:"msg" bson:"msg"`
	Time  string  `json:"time" bson:"time"`
}

type counter struct {
	Count float64 `bson:"seq_value"`
}

// albums slice to seed record album data.
var messages = []boardMessage{
	{Count: 1, Msg: "Good morning!", Time: "2022-01-30T08:01:46.356Z"},
	// {Count: 2, Msg: "Nyanpasu!", Time: "2022-01-30T08:03:10.862Z"},
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}

// postAlbums adds an album from JSON received in the request body.
func postMessages(c *gin.Context) {
	var newMessage boardMessage
	// Call BindJSON to bind the received JSON to
	// newMessage.
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	// MongoDB part
	client, ctx, cancel, err := connect(driver_URI)
	if err != nil {
		panic(err)
	}
	//  free resource when main function is returned
	defer close(client, ctx, cancel)

	returnMessage := counter{}
	collection := client.Database("nyanpasuSite").Collection("counters")
	err = collection.FindOne(
		ctx,
		bson.M{"_id": bson.M{"db": "nyanpasuSite", "coll": "messages"}},
		options.FindOne().SetProjection(bson.M{"_id": 0}),
	).Decode(&returnMessage)
	if err != nil {
		panic(err)
	}

	collection = client.Database("nyanpasuSite").Collection("messages")
	_, err = collection.InsertOne(ctx, newMessage)
	if err != nil {
		panic(err)
	}

	newMessage.Count = returnMessage.Count + 1
	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}

func main() {
	router := gin.Default()
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run("localhost:8080")
}
