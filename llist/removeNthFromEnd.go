package llist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{Next: head}
	first := head
	second := head

	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return p.Next
}
