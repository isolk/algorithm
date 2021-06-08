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

func NewCycle(values []int, pos int) *ListNode {
	head := &ListNode{}
	cur := head
	var tmp *ListNode = nil
	for i, value := range values {
		cur.Next = &ListNode{Val: value}
		cur = cur.Next
		if i == pos {
			tmp = cur
		}
	}
	cur.Next = tmp
	return head.Next
}

func Print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
