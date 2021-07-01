func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	
	stack := make([]*Node, 0)
	result := make([]int, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		result = append(result, node.Val)
	
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return result
}