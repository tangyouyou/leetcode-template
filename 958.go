package leetcode

//958. 二叉树的完全性检验
//给定一个二叉树，确定它是否是一个完全二叉树。

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	end := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			end = true
		} else {
			if end == true {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return true
}