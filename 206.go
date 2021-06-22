package leetcode

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev

		prev = head
		head = temp
	}

	return prev
}
























func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev

		prev = head
		head = temp
	}

	return prev
}