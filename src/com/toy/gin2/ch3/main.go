package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", startPage)
	r.Run("localhost:8000")
}

func startPage(c *gin.Context) {
	var person Person
	//POST请求时会首先检查请求体
	/*if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/
	if err := c.ShouldBindQuery(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Name":    person.Name,
		"Address": person.Address,
	})
}
