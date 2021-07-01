func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || m <= 0 || n <= 0 || m > n {
        return nil
    }   

    dummy := &ListNode{Val : 0}
    dummy.Next = head

    pre := dummy

    for i := 0; i < m-1; i++ {
        pre = pre.Next
    }

    right := pre
    for i := 0; i < n-m+1; i++ {
        right = right.Next
    }

    // 断开链表
    // 反转链表
    // 拼接链表
}