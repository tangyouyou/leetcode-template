package leetcode

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}

	return res
}
