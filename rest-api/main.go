package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"events": []string{"event1", "event2", "event3"},
		"other": map[string]map[string]string{
			"key1": {"subkey1": "value1", "subkey2": "value2"},
			"key2": {"subkeyA": "valueA", "subkeyB": "valueB"},
		},
	})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id": eventId,
	})
}
