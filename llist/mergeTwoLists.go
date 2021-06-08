package llist

// leetcode：21
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for {
		// 任意一个链表没有节点后，把另外一个链表直接放在后面，然后返回即可
		if l1 == nil {
			cur.Next = l2
			break
		}

		if l2 == nil {
			cur.Next = l1
			break
		}

		// 每次只比较一次，取值小的添加到新链表中
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		// 注意旧链表的指针移动后，新链表指针也要相应移动
		cur = cur.Next
	}
	return head.Next
}
