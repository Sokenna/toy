package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url+"_1", ch)
		go fetch(url+"_2", ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
	sec := time.Since(start).Seconds()
	fmt.Printf("sec:%.fs", sec)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	url2 := ""
	if strings.HasSuffix(url, "_1") {
		url2 = strings.TrimSuffix(url, "_1")
	}
	if strings.HasSuffix(url, "_2") {
		url2 = strings.TrimSuffix(url, "_2")
	}
	resp, err := http.Get(url2)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	dst, _ := os.Create(strings.ReplaceAll(strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://"), ".", "_") + ".html")
	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	dst.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs  %7d  %s", time.Since(start).Seconds(), nbytes, url)

}
