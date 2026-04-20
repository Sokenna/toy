package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		c.String(http.StatusOK, "pong")
	})
	router.Run(":8000")
}
