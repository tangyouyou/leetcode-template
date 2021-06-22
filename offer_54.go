package leetcode

//给定一棵二叉搜索树，请找出其中第k大的节点。
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[0:len(stack)-1]
		root = node.Right
	}

	return res[len(res) - k]
}