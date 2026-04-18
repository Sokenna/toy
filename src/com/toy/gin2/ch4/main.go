package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var login LoginForm
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"User": login.User, "Password": login.Password})
	})
	router.Run("localhost:8000")
}
