package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed templates
var tmplFS embed.FS

func main() {

	sub, _ := fs.Sub(tmplFS, "templates")
	println("=== 检查嵌入的文件 ===")
	err := fs.WalkDir(tmplFS, ".", func(path string, d fs.DirEntry, err error) error {
		println("文件:", path)
		return nil
	})
	if err != nil {
		println("遍历失败:", err.Error())
	}
	println("====================")

	router := gin.Default()
	router.LoadHTMLFS(http.FS(sub), "*.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "From FS"})
	})
	router.Run(":8000")
}
