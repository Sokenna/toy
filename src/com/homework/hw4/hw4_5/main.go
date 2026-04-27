package main

import "fmt"

func main() {
	s := []string{"H", "e", "l", "l", "o", "W", "o", "r", "l", "l", "l", "l", "d", "d"}
	s = removeRepeat(s)
	fmt.Println(s)
	s = []string{"H", "e"}
	s = removeRepeat(s)
	fmt.Println(s)
	s = []string{"H"}
	s = removeRepeat(s)
	fmt.Println(s)
	s = []string{"H", "o", "o", "o", "o", "o", "o"}
	s = removeRepeat(s)
	fmt.Println(s)
}

func removeRepeat(s []string) []string {
	ix := 0
	count := 0
	for ix < len(s)-count-1 {
		if s[ix] == s[ix+1] {
			copy(s[ix:], s[ix+1:])
			count++
		} else {
			ix++
		}
	}
	return s[:ix+1]
}
