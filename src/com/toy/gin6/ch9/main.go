package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	//router.SetTrustedProxies([]string{"192.168.1.2"})
	router.SetTrustedProxies(nil)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"client_ip": c.ClientIP(), "header": c.Request.Header, "IP": c.Request.Host})
	})
	router.Run(":8000")
}
