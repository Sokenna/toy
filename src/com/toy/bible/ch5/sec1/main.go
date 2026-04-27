package main

import "fmt"

// 函数声明包含 关键字func 函数名（形参列表 形参类型）(返回值列表 返回值类型) {函数体}
func main() {
	var n = 2

	fmt.Println(n)
	noChangeNum(n)
	fmt.Println(n)
	changeNum(&n)
	fmt.Println(n)
}
func changeNum(n *int) {
	*n = 10
}
func noChangeNum(n int) {
	n = 10
}
