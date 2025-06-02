package medium

func reorderList(head *ListNode) {
	//  1 2 3 4 5  //  middle = 3
	//  1 5 2 4 3
	mid := find_middle(head)
	reverse := reverse_list(mid.Next)
	mid.Next = nil
	func_head := head
	for func_head != nil && reverse != nil {
		temp1 := func_head
		temp2 := reverse
		func_head = func_head.Next
		reverse = reverse.Next

		temp1.Next = temp2
		temp2.Next = func_head
	}
}

func find_middle(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse_list(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	var first, second *ListNode = nil, head
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	return first
}
