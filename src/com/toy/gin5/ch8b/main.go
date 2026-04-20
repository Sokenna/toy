package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"

	"time"
)

// 使用zerolog

func ZeroLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Str("client_ip", c.ClientIP()).
			Msg("request")
	}
}
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if gin.Mode() == "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	}
	router := gin.New()
	router.Use(ZeroLogMiddleware())
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8000")
}
