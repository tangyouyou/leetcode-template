//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。 
// 注意：两结点之间的路径长度是以它们之间边的数目表示。 
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	iMax := 0

	depth(root, &iMax)

	return iMax
}

func depth(root *TreeNode, iMax *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, iMax)
	right := depth(root.Right, iMax)

	*iMax = max(*iMax, left + right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}