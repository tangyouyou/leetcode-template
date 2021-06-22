package leetcode

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (m *ListNode) String() string{
	str := ""
	for m != nil {
		if str == "" {
			str = fmt.Sprintf("%d", m.Val)
		} else {
			str = fmt.Sprintf("%s->%d", str, m.Val)
		}
		m = m.Next
	}
	return str
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
