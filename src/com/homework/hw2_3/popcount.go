package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println("7有多少个1：", PopCount1(7))
}

var pc [256]byte = func() (pc [256]byte) {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount1(x uint64) int {
	return int(func(x uint64) (r byte) {
		for i := 0; i < 8; i++ {
			r += pc[byte(x>>(i*8))]
		}
		return r
	}(x))
}

func PopCount2(x uint64) int {
	n := 0
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}
func PopCount3(x uint64) int {
	return bits.OnesCount(uint(x))
}
