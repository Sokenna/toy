package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"html": "<!DOCTYPE html><html lang=\"en\"><head>    <meta charset=\"UTF-8\">    <title>Title</title></head><body><b>Hello, world!</b></body>\n</html>"})
	})
	route.GET("/pureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{"html": "<!DOCTYPE html><html lang=\"en\"><head>    <meta charset=\"UTF-8\">    <title>Title</title></head><body><b>Hello, world!</b></body>\n</html>"})
	})
	route.Run(":8000")
}
