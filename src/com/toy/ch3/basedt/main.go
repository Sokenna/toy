package main

import (
	"fmt"
	"math"
)

// go数据类型分为四类1.基本类型2.引用类型3.复合类型4.接口类型
/*
1.基本类型
	1.1数字
		1.1.1 整型 有符合int int8 int16 int32 int64 无符号uint uint8 uint16 uint32 uint64  rune uintptr
		1.1.2 浮点类型 float32 float64
	1.2布尔型
	1.3字符串
	byte,char,int,int8,int32,int64,unint,unint8,

*/
var p uintptr

func main() {
	o := 0666
	fmt.Printf("%d %[1]o", o)
	fmt.Println(math.MaxFloat64)
	var f float32 = 1 << 24
	fmt.Println(f == f+1)
}
