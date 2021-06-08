package llist

// leetcode:141
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head
	for fast != slow {
		// 1. 快指针在移动时，需要检查是否走到了空节点，否则取孙子节点就会崩溃。
		// 2. 任意时候next有nil，肯定就没环了，可以直接返回。
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
