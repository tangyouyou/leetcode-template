package leetcode

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)

	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil  {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}