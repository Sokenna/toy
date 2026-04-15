package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "GET",
	})
}
func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "POST",
	})
}
func puting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "PUT",
	})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "PATCH",
	})
}
func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method": "DELETE",
	})
}
func options(c *gin.Context) {
	c.Status(http.StatusOK)
}
func head(c *gin.Context) {
	c.Status(http.StatusOK)
}
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/somePost", posting)
	router.PUT("/somePut", puting)
	router.PATCH("/somePatch", patching)
	router.DELETE("/someDelete", deleting)
	router.HEAD("/someHead", head)
	router.OPTIONS("someOption", options)
	err := router.Run("localhost:8000")
	if err != nil {
		return
	}
}
