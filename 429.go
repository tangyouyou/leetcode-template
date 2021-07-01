func levelOrder(root *Node) [][]int {
    if root == nil {
    	return [][]int{}
    }

    result := make([][]int, 0)
    queue := make([]*Node, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
    	l := len(queue)
    	ans := make([]int, 0)
    	for i := 0; i < l; i++ {
    		node := queue[i]
    		for j := 0; j < len(node.Children); j++ {
    			queue = append(queue, node.Children[j])
    		}
    		ans = append(ans, node.Val)
    	}

    	queue = queue[l:]
    	result = append(result, ans)
    }

    return result
}