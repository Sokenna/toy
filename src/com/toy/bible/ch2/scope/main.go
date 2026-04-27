package main

import "fmt"

func f() {}

var g = 10

func main() {
	f := 10
	fmt.Println(f) //local var shadows package-level func f
	fmt.Println(g) //package-level g
	//fmt.Println(h) //undefine h
	x := "hello!"
	fmt.Printf("local1 x:%p\n", &x)
	for i := 0; i < len(x); i++ {
		x := x[i]
		fmt.Printf("local2 x%p\n", &x)
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("local3 x%p\n", &x)
			fmt.Printf("%c\n", x)
		}
	}

}
