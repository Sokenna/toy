package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// cookie
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, false)
		}
		fmt.Printf("Cookie value:%s \n", cookie)
		c.SetCookie("gin_cookie", "test", -1, "/", "localhost", false, false)
	})
	router.Run(":8000")
}
