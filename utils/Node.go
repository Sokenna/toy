package utils

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func Travel(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	Travel(node.Left)
	Travel(node.Right)
}
