func demo_226(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left,root.Right = root.Right,root.Left
	demo_226(root.Left)
	demo_226(root.Right)

	return root
}