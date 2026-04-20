package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// autotls 暂时无法测试
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ping")
	})
	log.Fatal(autotls.Run(router, "example1.com"))
}
