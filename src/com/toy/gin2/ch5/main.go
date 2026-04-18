package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name" xml:"name" `
	Address  string    `form:"address" xml:"address"`
	Birthday time.Time `form:"birthday" xml:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person.Name, person.Address, person.Birthday)
	c.JSON(http.StatusOK, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
	})
}
func main() {
	r := gin.Default()
	r.GET("/testing", startPage)
	r.POST("/testing", startPage)
	r.Run("localhost:8000")
}
