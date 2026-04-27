package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	server := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 25,
	}
	log.Fatal(server.ListenAndServe())
}
