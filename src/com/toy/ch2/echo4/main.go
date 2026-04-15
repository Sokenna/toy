package main

import (
	"flag"
	"fmt"
	"github.com/Sokenna/toy/src/com/toy/ch2/tempconv"
	"strings"
)

// var 变量名称 类型 = 表达式  "类型"或"= 表达式"可以省略一个
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var v int

var c tempconv.Celsius
var f tempconv.Fahrenheit

func main() {
	c = tempconv.FToC(212.0)
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(c.String())
	fmt.Printf("%T\n", 0)
	fmt.Println(c == 0)
	fmt.Println(f == 0)
	//fmt.Println(c == f) //compile error:type mismatch

	fmt.Printf("%g\n", tempconv.BiolingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BiolingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	//fmt.Printf("%g\n", boilingF-tempconv.FreezingC)//compile error:type mismatch
	m := make(map[string]int)
	v, ok := m["hello"]
	fmt.Println("v=", v, ok)

	fmt.Println()
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	p := new(int)
	fmt.Println(*p)
	*p = 10
	fmt.Println(*p)
	fmt.Println(newAddr()) //return new addr every time(except nil)
	fmt.Println("-----")
	zeroV()
}
func newAddr() bool {
	p, q := new(int), new(int)
	return p == q
}
func zeroV() {
	p, q := new(struct{}), new(struct{})
	fmt.Println(&p, &q)
	var p1, q1 struct{}
	fmt.Printf("%p,%p\n", &p1, &q1)
	fmt.Println(&p1 == &q1)
	var p2, q2 [0]int //长度为零的数组
	fmt.Printf("%p,%p\n", &p2, &q2)
	fmt.Println(&p2 == &q2)
	fmt.Println("--------------------------------------------")
	fmt.Println("48和82最大公约数：", gcd(82, 48))
	fmt.Println("--------------------------------------------")
	fmt.Println("fib 5:", fib(10))
}
func incr(p *int) int {
	*p++
	return *p
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

/*
      0
    1   1
  1   2   1
1   3   3   1

*/

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
