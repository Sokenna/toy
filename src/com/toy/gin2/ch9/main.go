package main

import (
	"encoding"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)
var _ binding.BindUnmarshaler = (*Birthday)(nil)

func main() {

	r := gin.Default()

	var request struct {
		Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
		Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
		BirthdaysDefault []Birthday `form:"birthdayDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
	}
	r.GET("/test", func(c *gin.Context) {
		_ = c.BindQuery(&request)
		c.JSON(http.StatusOK, request)
	})
	var param struct {
		Birthday         Birthday   `form:"birthday"`
		Birthdays        []Birthday `form:"birthdays" collection_format:"csv"`
		BirthdaysDefault []Birthday `form:"birthdayDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
	}
	r.GET("/test2", func(c *gin.Context) {
		_ = c.BindQuery(&param)
		c.JSON(http.StatusOK, param)
	})
	r.Run("localhost:8000")
}
