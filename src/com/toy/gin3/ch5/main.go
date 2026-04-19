package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	route.GET("/user", func(c *gin.Context) {
		user := gin.H{"name": "Lena", "role": "admin"}
		switch c.Query("format") {
		case "xml":
			c.XML(http.StatusOK, user)
		case "yaml":
			c.YAML(http.StatusOK, user)
		default:
			c.JSON(http.StatusOK, user)

		}
	})
	route.Run(":8000")
}
