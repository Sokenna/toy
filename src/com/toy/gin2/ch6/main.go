package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Person struct {
	Name      string    `form:"name,default=William"`
	Age       int       `form:"age,default=10" `
	Friends   []string  `form:"friends,default=Will;Bill"`
	Addresses [2]string `form:"addresses,default=Foo Bar" collection_format:"ssv"`
	LapTimes  []int     `form:"lapTimes,default=1;2;3" collection_format:"multi"`
}

func main() {
	binding.EnableDecoderDisallowUnknownFields = true
	r := gin.Default()

	r.POST("/person", func(c *gin.Context) {
		var req Person
		if err := c.ShouldBind(&req); err != nil { // infers binder by Content-Type
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, req)
	})
	_ = r.Run(":8000")
}
