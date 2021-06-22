package leetcode

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

/**
* Definition for singly-linked list.
* type ListNode struct {
* Val int
* Next *ListNode
* }
 */
func ReverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = ReverseKGroup(node, k)

	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := first

	for first != last {
		temp := first.Next
		first.Next = prev

		prev = first
		first = temp
	}

	return prev
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = ReverseKGroup(node, k)

	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := first
	for first != last {
		temp := first.Next
		first.Next = prev

		prev = first
		first = temp
	}

	return prev
}
