package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

var pc [256]byte

//arr 1.固定长度、特定类型元素、零个或多个元素组成，长度固定
//[...]int{1,2,3}在长度位置出现... 表示数组的长度根据初始化值的数量来确定

// 初始化表
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, i2 := range a {
		fmt.Printf("index:%d data:%d \n", i, i2)
	}
	for _, v := range a {
		fmt.Printf("data:%d\n", v)
	}
	var q = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	for _, v := range q {
		fmt.Printf("data:%d\n", v)
	}
	for _, v := range r {
		fmt.Printf("data:%d\n", v)
	}
	var k = [...]int{4, 5, 6}
	for _, v := range k {
		fmt.Printf("data:%d\n", v)
	}

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println("symbol len:", len(symbol))
	s1 := "x"
	s2 := "X"
	s1256 := sha256.Sum256([]byte(s1))
	s2256 := sha256.Sum256([]byte(s2))

	fmt.Printf("%x \n%x \n", s1256, s2256)
	var x uint64 = 0x123456789ABCDEF0
	fmt.Printf("%08b \n%08b \n%08b\n", x, x>>(0*8), byte(x>>(0*8)))
	fmt.Printf("%08b \n%08b \n%08b\n", x, x>>(1*8), byte(x>>(1*8)))
	fmt.Printf("%08b \n%08b \n%08b\n", x, x>>(2*8), byte(x>>(2*8)))

	r1 := [...]int{99: -1}
	fmt.Println(r1)
	a32 := [32]int{1, 2, 3, 4, 5}
	fmt.Println("原始数组:", a32)
	zero(&a32)
	fmt.Println("清零后:", a32)
}

func zero(ptr *[32]int) {
	for i := 0; i < 32; i++ {
		ptr[i] = 0
	}
}
