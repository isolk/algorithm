package llist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(values []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, value := range values {
		cur.Next = &ListNode{Val: value}
		cur = cur.Next
	}
	return head.Next
}

func Print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
