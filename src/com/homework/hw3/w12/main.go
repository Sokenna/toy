package main

import (
	"fmt"
)

func main() {
	s1 := "hellohh"
	s2 := "hlloole"
	fmt.Println(isDWEO(s1, s2))
}

func isDWEO(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range s1 {
		count[c]++
	}
	for _, c := range s2 {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}
	return true
}
