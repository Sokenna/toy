package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("authRequired...")
		name := c.PostForm("name")
		password := c.PostForm("password")
		fmt.Println(name, password)
		if password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码不能为空！"})
			c.Redirect(http.StatusUnauthorized, "/api/health")
		}
		c.Next()

	}
}
func main() {
	router := gin.Default()
	//Public routes -- no auth required
	public := router.Group("/api")
	{
		public.GET("/health", healthCheck)
	}
	private := router.Group("/api")
	private.Use(authRequired())
	{
		private.GET("/profile", getProfile)
		private.POST("/setttings", updateSetttings)
	}
	err := router.Run("localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stdout, "err:%v", err)
		return
	}
}

func getProfile(c *gin.Context) {
	fmt.Println("getProfile...")
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func updateSetttings(c *gin.Context) {
	fmt.Println("updateSetttings...")
	c.Status(http.StatusOK)
}
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
