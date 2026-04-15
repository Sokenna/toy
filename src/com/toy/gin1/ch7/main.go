package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/old", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/result")
	})
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	})

	router.GET("/final", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})
	router.GET("/result", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "redirect here!"})
	})

	router.Run("localhost:8000")
}
