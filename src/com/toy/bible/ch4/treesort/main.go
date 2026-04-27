package main

import "fmt"

// 结构体的零食值是每个成员都是零值
func main() {
	p := Point{
		X: 1,
		Y: 2,
	}
	p = scale(p, 5)
	fmt.Println(p)
}

func scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

type Point struct {
	X int
	Y int
}
type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values, root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, t.value)
	} else {
		t.right = add(t.right, t.value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
