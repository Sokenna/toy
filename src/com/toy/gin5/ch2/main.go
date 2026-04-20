package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

func main() {
	router := gin.New()

	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(&lumberjack.Logger{
		Filename:   "toy.log",
		MaxSize:    8,
		MaxAge:     25,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}, os.Stdout)
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s  %s  %s  %d  %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage)
	}))
	/*gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "toy.log",
		MaxSize:    8,
		MaxAge:     25,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}*/
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8000")
}
