package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hey", "status": http.StatusOK})
	})
	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hey", "status": http.StatusOK})
	})
	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "Hey", "status": http.StatusOK})
	})

	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		lable := "test"
		data := &protoexample.Test{
			Label:         &lable,
			Type:          nil,
			Reps:          reps,
			Optionalgroup: nil,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	router.Run(":8000")
}
