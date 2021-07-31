package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func MidTraverse(node *Node){
	if node == nil{
		return
	}
	MidTraverse(node.Left)
	fmt.Println(node.Value)
	MidTraverse(node.Right)
}
