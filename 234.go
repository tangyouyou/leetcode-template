// 判断是否回文链表；思路：找到中间节点，断开链表，并对尾部链表进行反转，最后判断链表的值是否相等
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    tail := reverse(slow.Next)

    // 断开头部的链表
    slow.Next = nil
    for head != nil && tail != nil {
        if head.Val != tail.Val {
            return false
        }
        head = head.Next
        tail = tail.Next
    }

    return true
}

func reverse(head *ListNode) *ListNode {
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