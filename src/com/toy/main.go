package main

import (
	"fmt"
	"github.com/Sokenna/toy/src/com/toy/leetcode"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type I interface {
	M()
}
type User struct {
	Name    string
	Age     int
	Address string
}

func (u User) String() string {
	return fmt.Sprintf("Name:%s , Age:%d , Address:%s", u.Name, u.Age, u.Address)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v , %s", e.When, e.What)
}
func run() error {
	var m = MyError{
		When: time.Now(),
		What: "it did't work",
	}
	if m.What != "" {
		return error(&m)
	}
	return nil
}

type MyReader struct {
}

func (receiver MyReader) Reader(v []byte) (int, error) {
	return len(v), nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	switch {
	case x >= 65 && x <= 77:
		fallthrough
	case x >= 97 && x <= 109:
		x = x + 13
	case x >= 78 && x <= 90:
		fallthrough
	case x >= 110 && x <= 122:
		x = x - 13
	}
	return x
}
func (r rot13Reader) Read(b []byte) (int, error) {
	n, _ := r.r.Read(b)
	for key, value := range b {
		b[key] = rot13(value)
	}
	return n, nil
}
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func (t *List[T]) add(v T) {
	node := &List[T]{val: v}
	t.next = node
}
func (t *List[T]) String() string {
	var n = t
	var s []T
	for n.next != nil {
		s = append(s, n.val)
		n = n.next
	}
	return fmt.Sprintf("%v", s)
}
func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func fabonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

type SafeCounter struct {
	mu sync.RWMutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var count = make(map[string]int)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	if v, ok := count[url]; ok {
		count[url] = v + 1
	} else {
		if url != "" {
			count[url] = 1
		}
	}
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type Jbody struct {
	FileName string `json:"file_name"`
	FileId   int    `json:"file_id"`
	Tag      string `json:"tag"`
	Gender   int    `json:"gender"`
	Hello    string `json:"hello"`
}

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	leetcode.ExcuteFunc()
	gin.Default()
	/*prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)*/
}
func spinner(delay time.Duration, c chan bool) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
	c <- true
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func Resolve() {
	str := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAtAZPPiQdEyNbPldm+obngxkgaU9iO3RRS0m3Bhv+AiwBtuWx
icsIQM/3RO/WH7Tojx7HN/k5ojcof7S5trmqmBd27Lww/11lYaJhvN1gJmobOT/2
Z51M0INHZxWQOvYPB7hqYMoHB01eyjILT6620YvVFdUjkLlvYL4Hf/Dr/CWMJAb0
PflGWlbeCR5tY4fVG2sBRF2yLtGOziPMdPfNB6/UxbvGmHJTDCgRg5QwnsiiFyFq
rHTTsFpxtB5E4GZsVOOX/e05wc95ACRZyyIZByUlgfxRRsH9rSoHamTDzGdXVJHr
rVWDGWjnwh4ViJAr1ClG+TQDG4ddky8OdTaJGwIDAQABAoIBAE0R+nDXEx9zXoe0
8WBFyu1kH20WUC614GvD9jgGWdtiLaW4diZQporgvbJknx3Z3EoRWiaLavIYfR0X
KP6iQ/dy7DXzZ2KcAGApHYC2oYgwmtEKohm/zPfji/kE0Ud8ufiLKfef5hxpMjCL
9i251YimPaZoJh66VIVtNVa/tLewPqa8IGIairA+U7r2bK0YHvnHXdiwguvNH/v6
CnOzW+eC/EPb4ym1WGD7gzBY22213BYwXuPH12cerKjz6thHkCmwGNKel2J/ByPA
czflFIlZc2HD77koMvX7+qqXsCkktGfO7iW2k3KQ498bUbSj1cRnDjAmqoo3cjim
YLob/3kCgYEA7blRgVsNBeWqGEUpfGiEBF7Bq8Z5Rju32ExZjeekRbLxukr32Bs1
RcEcY5oMMP1P08Rj43RbC6guSf0zWs6lS0uw0ahzO2zKznIJmFZeFH9H0Zx0dEVD
1R2z5a+JleEYvpARJWinmOxlIeOE4wAdkZ+b3QBL8HIDZBqwsF2/WOUCgYEAwd1n
LfpMyeDoMKm0b5pDStkgSBHGf2GosUAfULPedWTUZ7x8df0OUQm3fO8ybt5rYs+z
zQ4V+ZFo/R3clPl4kZc+KD1y0l7ZtmHWVT0thudXVw+N49bNDhbUcYs+D7zR7Tn0
9FJgnxmWB/vuZtMbUyuqvj4GOOveEEx7U2JqOf8CgYEAy529yDR61SBALwWeYScM
XMnkzQL1AUlJUQUkd71/IQwCrHRmET5MxPCBJnbeFmACLfq7LCmPik6FR26OvCrf
vbyZORs1iCLJG0bmHQttDdVvaIS/4o8paLNGUWenSy9AiIjRht87HmrsvLgqLFz+
9qplUEhewxaE/Qs3Z/kEWlkCgYBUpJxWIMDCK+Z0yMO1Ln9PALntfYVTPIpwyXRi
wgXQ77Mlj9Avm8tSapGohK9aZNYyEEmKsCm8C1bxnMYgbUPNHWHQI/QsEamwzcrM
8KceDYe+xVXIMpLZIfKjmI4CRQjLMNDYk9cH4B33YnHWhzmY+KnV4jLJS4JrT1lR
MNj8dQKBgQDhAG42zG3rdjH1QNUTUPvQ4dvSgXPc4xyELuquVGlXKR3n1L2wDSON
/ow+vr9T9RTmxGgpf9Oa2mfcpuUD59TlGsXtnahUAn7fXaBRvRZTJLLKwADHM2If
EdGWkEH8ALxTjOnQgoGBoOWIs2fJQEczR6l5E8NfmpuoPZlmqzDzkQ==
-----END RSA PRIVATE KEY-----`
	key, err := ssh.ParsePrivateKey([]byte(str))
	if err != nil {
		return
	}
	fmt.Println(key)
}
func theMaximumAchievableX(num int, t int) int {
	return num + (2 * t)
}
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
func concat(c chan string) {
	c <- "hello"
	c <- "world"
	close(c)
}
func printChan(ch chan string) {
	for s := range ch {
		fmt.Println(s)
	}

}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
