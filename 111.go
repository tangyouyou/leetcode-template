package leetcode

import "math"

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}