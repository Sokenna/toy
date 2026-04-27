package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	f := "/Users/dsy/ch3/files/test.txt"
	fmt.Println(basename(f))
	fmt.Println(filepath.Base(f))
}

func basename1(s string) string {
	p := strings.LastIndex(s, "/")
	e := strings.LastIndex(s, ".")
	return s[p+1 : e]
}
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
