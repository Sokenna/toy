package main

import "fmt"

func main() {

	fmt.Println(PopCountBetter(7))
	{

	}
}

func PopCountBetter(x uint64) int {
	count := 0
	for x != 0 {
		fmt.Printf("%08b   %08b\n", x, x-1)
		x &= x - 1
		count++
	}
	return count
}
