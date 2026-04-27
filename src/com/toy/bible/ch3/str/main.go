package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s = "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); i++ {
		r, s := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c  ", r)
		i += s - 1 // 这里的i++是多余的，因为DecodeRuneInString已经返回了字符所占的字节数
	}
	fmt.Println()
	for _, i2 := range s {
		fmt.Printf("%c  ", i2)
	}
	fmt.Println()
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println(string(1234567))
	fmt.Println(strconv.FormatBool(true))
	if parseBool, err := strconv.ParseBool("true"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parseBool)
	}
	str1 := 'a'
	ustr := unicode.ToUpper(str1)
	fmt.Printf("%c\n", ustr)
}
