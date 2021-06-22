package leetcode

import (
	"fmt"
	"strconv"
)

// 二叉树的路径
func SumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result [][]int
	var arr = make([]int, 0)
	dfs1(root, arr, &result)

	num := 0
	for i := 0; i < len(result); i++ {
		nums := result[i]
		str := ""
		for _, v := range nums {
			str = fmt.Sprintf("%s%d", str,v)
		}
		count,_ := strconv.Atoi(str)
		num = num + count
	}

	return num
}

func dfs1(root *TreeNode, arr []int, result *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)

	// 根节点且值等于最后一个路径的值
	if root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*result = append(*result, tmp)
	}

	dfs1(root.Left,  arr, result)
	dfs1(root.Right, arr, result)
	arr = arr[0 : len(arr) - 1]
}