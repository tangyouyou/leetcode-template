package leetcode

func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	path := []int{}
	dfs131(&ret, root, path, sum)
	return ret
}

func dfs131(ret *[][]int, root *TreeNode, path []int, target int) {
	switch {
	case root == nil:
		return
	case root.Left == nil && root.Right == nil && root.Val == target:
		dst := make([]int, len(path)+1)
		copy(dst, append(path, root.Val))
		*ret = append(*ret, dst)
		return
	}
	path = append(path, root.Val)
	dfs131(ret, root.Left, path, target-root.Val)
	dfs131(ret, root.Right, path, target-root.Val)
}
