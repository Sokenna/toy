package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

// 运行多个服务

var g errgroup.Group

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Welcome to router01"})
	})
	return e
}
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Welcome to router02"})
	})
	return e
}
func main() {
	server01 := &http.Server{
		Addr:           ":8000",
		Handler:        router01(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server02 := &http.Server{
		Addr:           ":8001",
		Handler:        router02(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
