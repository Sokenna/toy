package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/codec/json"
	_ "github.com/json-iterator/go"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

var customConfig = jsoniter.Config{
	EscapeHTML:             false, //true {"html":"\u003ch1\u003e\u003cbr\u003eHello World!\u003cbr\u003e\u003ch1\u003e","name":"sean","title":"custom json"}
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

type customJsonApi struct{}

func (j customJsonApi) Marshal(v any) ([]byte, error) {
	return customConfig.Marshal(v)
}

func (j customJsonApi) Unmarshal(data []byte, v any) error {
	return customConfig.Unmarshal(data, v)
}
func (j customJsonApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return customConfig.MarshalIndent(v, prefix, indent)
}

func (j customJsonApi) NewEncoder(writer io.Writer) json.Encoder {
	return customConfig.NewEncoder(writer)
}

func (j customJsonApi) NewDecoder(reader io.Reader) json.Decoder {
	return customConfig.NewDecoder(reader)
}

// 运行时自定义编解码器
func main() {
	json.API = customJsonApi{}
	router := gin.Default()
	router.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "sean",
			"title": "custom json",
			"html":  "<h1><br>Hello World!<br><h1>",
		}
		c.JSON(http.StatusOK, data)
	})
	router.Run(":8000")
}
