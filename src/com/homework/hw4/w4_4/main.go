package main

import "fmt"

func main() {
	a := []int{0, 1}
	a = rotate(a, 1)
	fmt.Println(a)
}
func rotate(s []int, x int) []int {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[len(s)-i-1]
		s[i], s[len(s)-i-1] = tmp, s[i]
	}
	return s
}
