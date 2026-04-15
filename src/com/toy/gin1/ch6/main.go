package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			users := v1.Group("/users")
			users.GET("/", listUsers)
			users.GET("/:id", getUser)

			posts := v1.Group("/posts")
			posts.GET("/", listPosts)
			posts.GET("/:id", getPost)
		}
	}
	router.Run("localhost:8000")
}

func listUsers(c *gin.Context) {
	c.String(http.StatusOK, "listUsers")
}

func getUser(c *gin.Context) {
	c.String(http.StatusOK, "getUser")
}

func listPosts(c *gin.Context) {
	c.String(http.StatusOK, "posts")
}
func getPost(c *gin.Context) {
	c.String(http.StatusOK, "getPost")
}
