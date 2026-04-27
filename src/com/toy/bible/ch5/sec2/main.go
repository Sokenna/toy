package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/spider", func(c *gin.Context) {
		response, err := http.Get("https://www.runoob.com/try/ajax/demo_htmldom.htm")
		if err != nil {
			log.Printf("request err: %v", err)
			return
		}
		reader := response.Body
		doc, err := html.Parse(reader)
		if err != nil {
			return
		}
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	})
	router.Run(":8000")
}
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
