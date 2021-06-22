package leetcode

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val:0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}

	return dummy.Next
}





































func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val:0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
			head = head.Next
		} else {
			head.Next = l2
			l2 = l2.Next
			head = head.Next
		}
	}

	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}

	return dummy.Next
}
