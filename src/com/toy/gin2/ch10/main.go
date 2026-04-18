package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type testHeader struct {
	Rate          int    `header:"rate"`
	Domain        string `header:"domain"`
	Authorization string `header:"authorization" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		h := testHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain, "message": "通过"})
	})
	router.Run(":8000")
}
