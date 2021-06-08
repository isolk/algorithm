package llist

// leetcode:206
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTmp
	}
	return pre
}
