package llist

// leetcode:141
func HasCycle(head *ListNode) bool {
	fast := head.Next
	slow := head
	for slow != nil && fast != nil && fast != slow {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	return fast == slow
}
