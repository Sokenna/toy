package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/local/file", func(c *gin.Context) {
		c.File("local/test.go")
	})
	router.GET("/fs/file", func(c *gin.Context) {
		c.FileFromFS("/fs/test.go", http.Dir("my_file_system"))
	})
	router.GET("/download", func(c *gin.Context) {
		c.FileAttachment("local/report.xlsx", "report1.xlsx")
	})
	router.Run(":8000")
}
