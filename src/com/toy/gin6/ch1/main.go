package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 自定义http设置
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "pong")
	})
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
