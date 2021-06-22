package leetcode

//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeft := true

	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, l)

		for i := 0; i < l; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeft {
				ans[i] = node.Val
			} else {
				ans[l - i -1] = node.Val
			}
		}
		isLeft = !isLeft
		result = append(result, ans)
		queue = queue[l:]
	}

	return result
}
































func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeft := true
	result := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeft {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}
		}
		result = append(result, ans)
		isLeft = !isLeft
		queue = queue[l:]
	}

	return result
}