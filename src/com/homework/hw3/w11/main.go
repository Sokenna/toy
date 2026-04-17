package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))              // 123
	fmt.Println(comma("12345"))            // 12,345
	fmt.Println(comma("1234567"))          // 1,234,567
	fmt.Println(comma("1234567.89"))       // 1,234,567.89
	fmt.Println(comma("-1234567.89"))      // -1,234,567.89
	fmt.Println(comma("+987654321.12345")) // +987,654,321.12345
	fmt.Println(comma("0.1234"))           // 0.1234
}
func comma(s string) string {
	var buf bytes.Buffer
	sign := ""
	start := 0
	//处理符号
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		sign = s[:1]
		start = 1
	}
	buf.WriteString(sign)
	numPart := s[start:]
	dotIndex := bytes.Index([]byte(numPart), []byte("."))
	intPart := numPart
	fracPart := ""
	if dotIndex != -1 {
		intPart = numPart[:dotIndex]
		fracPart = numPart[dotIndex:]
	}

	n := len(intPart)

	if n <= 3 {
		buf.WriteString(intPart)
	} else {

		remain := n % 3
		if remain != 0 {
			buf.WriteString(intPart[:remain])
			buf.WriteString(",")
		}

		for i := remain; i < n; i += 3 {
			buf.WriteString(intPart[i : i+3])
			if i+3 < n {
				buf.WriteString(",")
			}
		}
	}
	buf.WriteString(fracPart)
	return buf.String()
}
