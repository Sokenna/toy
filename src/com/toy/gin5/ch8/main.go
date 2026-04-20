package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"time"
)

// 结构化日志log/slog
func RequestIDMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
func SlogMiddlewareHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		requestID, _ := c.Get("request_id")
		logger.Info("request",
			slog.String("request_id", requestID.(string)),
			slog.String("method", c.Request.Method),
			slog.String("path", path),
			slog.String("query", query),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("latency", time.Since(start)),
			slog.String("client_ip", c.ClientIP()),
			slog.Int("size", c.Writer.Size()),
		)
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				slog.Error("request error", slog.String("error", err.Error()))
			}
		}

	}
}
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := gin.New()
	router.Use(RequestIDMiddleWare())
	router.Use(SlogMiddlewareHandler(logger))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run(":8000")
}
