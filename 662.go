package leetcode

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, &TreeNode{0, root.Left, root.Right})

	for len(queue) > 0 {
		var left,right *int
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				newVal := node.Val * 2
				queue = append(queue, &TreeNode{newVal, node.Left.Left, node.Left.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}

			if node.Right != nil {
				newVal := node.Val * 2 + 1
				queue = append(queue, &TreeNode{newVal, node.Right.Left, node.Right.Right})
				if left == nil || *left > newVal {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
			}

			switch {
			case left != nil && right == nil, left == nil && right != nil:
				res = max(res, 1)
			case left != nil && right != nil:
				res = max(res, *right - *left  +1)
			}
		}
	}

	return res
}

