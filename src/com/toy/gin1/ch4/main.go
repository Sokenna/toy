package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

const (
	MaxUploadSize = 1 << 20
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

	router.POST("/upload_m", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fiels := form.File["files"]
		for _, file := range fiels {
			log.Println(file.Filename)
			dst := filepath.Join("./files/", filepath.Base(file.Filename))
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				log.Println(err)
				continue
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(fiels)))

	})
	router.POST("/upload_limit", uploadHandler)
	router.Run("localhost:8000")
}

func uploadHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)
	if err := c.Request.ParseMultipartForm(MaxUploadSize); err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": fmt.Sprintf("file too large (max:%d bytes)", MaxUploadSize)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, f, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	dst := filepath.Join("./files/", filepath.Base(f.Filename))
	err = c.SaveUploadedFile(f, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
