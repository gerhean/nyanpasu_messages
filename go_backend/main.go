package main

// https://go.dev/doc/tutorial/web-service-gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	router := gin.Default()
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run("localhost:8080")
}
