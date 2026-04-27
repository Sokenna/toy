package main

import "fmt" //import告诉编译器，需要哪些包

// 编译型语言，go工具链将源代码及其依赖转换成计算机指令（静态编译）
// 原生支持Unicode
// main包定义可独立执行的程序
func main() {
	fmt.Println("hello world!")
}
