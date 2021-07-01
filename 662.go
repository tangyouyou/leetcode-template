func demo_662(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, &TreeNode{0, root.Left, root.Right})

	ans := 0
	for len(queue) > 0 {
		l := len(queue)

		var left *int
		var right *int
		for i := 0; i < l; i++ {
			node := queue[i]
			newVal := 0
			if node.Left != nil {
				newVal = node.Val * 2
				queue = append(queue, &TreeNode{newVal, node.Left.Left, node.Left.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
			if node.Right != nil {
				newVal = node.Val * 2 + 1
				queue = append(queue, &TreeNode{newVal, node.Right.Left, node.Right.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}
		}
		queue = queue[l:]
		if (left == nil && right != nil) || left != nil && right == nil {
			ans = max(ans, 1)
		} else if left != nil && right != nil {
			ans = max(ans, *right - *left + 1)
		}
	}

	return ans
}