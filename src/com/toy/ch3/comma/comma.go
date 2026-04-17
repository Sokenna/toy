package main

import "fmt"

func main() {
	num := "10000000000"
	fmt.Println(comma(num))
	s := "abc"
	fmt.Printf("%s\n", s)
	b := []byte(s)
	fmt.Printf("%s\n", b)
	b[1] = 'd'
	fmt.Printf("%s\n", b)
	s = string(b)
	fmt.Printf("%s\n", s)

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
