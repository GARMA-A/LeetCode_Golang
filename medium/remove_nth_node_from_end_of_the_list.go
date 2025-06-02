package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	size := 0

	for ptr != nil { // O(n)
		ptr = ptr.Next
		size++
	}
	size = size - n
	if size == 0 {
		return head.Next
	}
	// half of the problem solved!
	cur := head
	for size > 1 {
		cur = cur.Next
		size--
	}
	cur.Next = cur.Next.Next
	return head
}
