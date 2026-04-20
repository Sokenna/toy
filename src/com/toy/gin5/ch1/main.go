package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

func main() {
	gin.DisableConsoleColor()
	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "toy.log",
		MaxSize:    100,
		MaxAge:     28,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8000")
}
