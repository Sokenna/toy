package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	loggerConfig := gin.LoggerConfig{
		SkipQueryString: false,
	}
	router.Use(gin.LoggerWithConfig(loggerConfig))
	router.Use(gin.Recovery())

	router.GET("/search", func(c *gin.Context) {
		q := c.Query("q")
		c.String(http.StatusOK, "searching for"+q)
	})
	router.Run(":8000")
}
