package main

import (
	"bufio"
	"fmt"
	"os"
)

// %d 十进制的数
// %x 16进制 %o 八进制  %b二进制
// %f %g %e 浮点数： 3.141593 3.141592653589793 3.141593e+00
// %s字符串
// %t 布尔值
// %c字符（rune）unicode码点
// %q 带双引号的字符串或单单引号的字符
// %v变量的自然形式
// %T变量的类型

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		} else {
			counts[input.Text()]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
