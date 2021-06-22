package leetcode

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == l -1 {
				result = append(result, node.Val)
			}
		}
		queue = queue[l:]
	}

	return result
}
