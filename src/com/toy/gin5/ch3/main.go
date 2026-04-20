package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	loggerConfig := gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}
	loggerConfig.Skip = func(c *gin.Context) bool {
		return c.Writer.Status() < http.StatusInternalServerError
	}
	router.Use(gin.LoggerWithConfig(loggerConfig))

	router.Use(gin.Recovery())
	router.GET("/metrics", func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/data", func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	})
	router.Run(":8000")
}
