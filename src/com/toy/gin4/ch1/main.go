package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Println(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//router.Use(gin.BasicAuth(gin.Accounts{"name": "admin", "password": "123456"}))
	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})
	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"type": "Login"})
		})
		authorized.GET("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"type": "Submit"})
		})
		authorized.GET("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"type": "Read"})
		})
	}
	{
		testing := authorized.Group("/testing")
		testing.Use(gin.BasicAuth(gin.Accounts{"name": "admin", "password": "123"}))
		testing.GET("/analytics", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"type": "Analytics"})
		})
	}
	router.Run(":8000")

}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
func benchEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}
