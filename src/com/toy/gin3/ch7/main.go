package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	routrer := gin.Default()
	routrer.Static("/assets", "./assets")
	routrer.Run(":8000")
}
