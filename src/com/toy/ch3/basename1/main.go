package main

import "fmt"

func main() {
	f := "/Users/dsy/ch3/files/test.txt"
	fmt.Println(basename(f))
}

/*
从后往前遍历，遇到第一个'/'截取之后的，再一次从后往前遍历，遇到'.'截取之前的。剩下的就是basename
*/
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			fmt.Printf("% c\n", s[i])
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
