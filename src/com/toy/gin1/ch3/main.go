package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Tang")
		lastName := c.Query("lastName")
		c.String(http.StatusOK, "Welcome %s %s", lastName, firstName)
	})
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,World!")
	})
	router.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON(http.StatusCreated, gin.H{"user": name})
	})
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": message,
			"nick":    nick,
		})
	})
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		/*c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})*/
	})

	router.Run("localhost:8000")
}
