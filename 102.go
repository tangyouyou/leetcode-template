package leetcode

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			ans = append(ans, node.Val)
		}
		queue = queue[l:]
		result = append(result, ans)
	}

	return result
}