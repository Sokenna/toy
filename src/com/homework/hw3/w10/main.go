package main

import (
	"bytes"
	"fmt"
)

func main() {
	num := "1000000"
	fmt.Println(comma(num))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	remain := n % 3
	if remain != 0 {
		buf.WriteString(s[:remain])
		buf.WriteString(",")
	}
	for i := remain; i < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
