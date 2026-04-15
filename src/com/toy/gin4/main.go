package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v  names: %v", ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println(file.Header)
		dst := filepath.Join("./files/", filepath.Base(file.Filename))
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			log.Println("保存文件时出错:", err)
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run("localhost:8000")
}
