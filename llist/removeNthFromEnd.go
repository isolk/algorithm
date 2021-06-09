package llist

// leetcode 19
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 先将first移动n个节点
	first := head
	for i := 0; i < n; i++ {
		first = first.Next
	}

	// 创建哑节点，然后往后移动，到最后的时候，second指向要删除节点的前一个节点
	tmp := &ListNode{Next: head}
	second := tmp
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 可能链表长度为0
	if second.Next != nil {
		second.Next = second.Next.Next
	}
	return tmp.Next
}
