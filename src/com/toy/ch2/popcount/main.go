package main

import "fmt"

var pc [256]byte = func() (pc [256]byte) {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func main() {
	fmt.Printf("%08b\n", 7)
	fmt.Printf("%08b\n", 7>>1)
	fmt.Println(PopCount(7))
}

/*func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}*/

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
