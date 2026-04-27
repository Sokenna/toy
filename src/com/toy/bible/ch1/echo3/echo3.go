package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	count := 100000
	args := make([]string, count)
	for i := 0; i < count; i++ {
		args[i] = "test"
	}
	os.Args = args
	Echo1()
	Echo2()
}

func echo3() {

	for i, arg := range os.Args[1:] {
		fmt.Printf("%d:%s \n", i, arg)
	}

}

// Echo1 4435ms
func Echo1() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	ms := time.Since(start).Milliseconds()
	fmt.Println(ms)
}

// Echo2 0
func Echo2() {
	start := time.Now()

	strings.Join(os.Args[1:], " ")
	ms := time.Since(start).Milliseconds()
	fmt.Println(ms)
}
