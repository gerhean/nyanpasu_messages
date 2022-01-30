package main

// https://go.dev/doc/tutorial/web-service-gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type boardMessage struct {
	Count float64 `json:"count,omitempty" bson:"count,omitempty"`
	Msg   string  `json:"msg" bson:"msg"`
	Time  string  `json:"time" bson:"time"`
}

type counter struct {
	Count float64 `bson:"seq_value"`
}

// An example being:
// {{Count: 1, Msg: "Good morning!", Time: "2022-01-30T08:01:46.356Z"}, }
var messages []boardMessage

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

func getMessages(c *gin.Context) {
	if len(messages) > 5 {
		c.IndentedJSON(http.StatusOK, messages[len(messages)-5:])
	} else {
		c.IndentedJSON(http.StatusOK, messages)
	}
}

func fetchMessagesFromDb() {
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
	count := returnMessage.Count

	collection = client.Database("nyanpasuSite").Collection("messages")
	filter := bson.M{"count": bson.M{"$gt": count - 5}}
	cursor, err := collection.Find(ctx,
		filter,
		options.Find().SetProjection(bson.M{"_id": 0}),
	)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &messages); err != nil {
		panic(err)
	}
}

// postAlbums adds an album from JSON received in the request body.
func postMessages(c *gin.Context) {
	var newMessage boardMessage
	// Call BindJSON to bind the received JSON to
	// newMessage.
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

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
	if returnMessage.Count > 1000 {
		panic("Too many entries in database!!!")
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// load .env file from given path
	_ = godotenv.Load(".env")
	// ignore error
	driver_URI = os.Getenv("MONGO_URI")

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	fetchMessagesFromDb()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run(":" + port)
}
