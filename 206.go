func demo_206(head *LinkedNode) *LinkedNode {
	if head == nil {
		return nil
	}

	var prev *LinkedNode

	for head != nil {
		temp := head.Next
		head.Next := prev

		prev = head
		head = temp
	}

	return prev
}