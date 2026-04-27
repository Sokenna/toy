package main

import "fmt"

// map 零值是nil
// map中的元素不是变量，不能进行取址操作
func main() {
	m := make(map[string]int)
	m["h"] = 1
	fmt.Println(m)
	for name, age := range m {
		fmt.Printf("%s,%d\n", name, age)
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
