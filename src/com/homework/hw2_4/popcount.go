package main

import "fmt"

func main() {
	fmt.Println(PopCountShift(7))
}
func PopCountShift(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		n += int(x & 1)
		x >>= 1
	}
	return n
}
