/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func getLength(head *ListNode) int {
    if head == nil {
        return 0
    }
    length := 0
    for head != nil {
        length++
        head = head.Next
    }

    return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {
        return nil
    }

    dummy := &ListNode{Val : 0}
    dummy.Next = head

    length := getLength(head)
    pre := dummy
    for i := 0; i < length - n; i++ {
        pre = pre.Next
    }
    pre.Next = pre.Next.Next

    return dummy.Next
}