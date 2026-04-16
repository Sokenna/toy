package main

import "fmt"

func f() int {
	return 10
}
func g(x int) int {
	return x * 2
}
func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y)
}
