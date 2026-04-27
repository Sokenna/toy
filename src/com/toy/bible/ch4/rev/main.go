package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(a[:])
	//fmt.Println(a)
	reverse(a[:2]) //1,0,2,3,4,5
	fmt.Println(a)
	reverse(a[2:]) //1,0,5,4,3,2
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	var s []int
	fmt.Println(len(s), s == nil)
	s = nil
	fmt.Println(len(s), s == nil)
	s = []int(nil)
	fmt.Println(len(s), s == nil)
	s = []int{}
	fmt.Println(len(s), s == nil)

	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
