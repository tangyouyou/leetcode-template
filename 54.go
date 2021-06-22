package leetcode

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	if rows == 0 || cols == 0 {
		return []int{}
	}
	left := 0
	right := cols -1
	top := 0
	bottom := rows -1
	result := make([]int, rows * cols)
	index := 0
	for left <= right && top <= bottom {
		for col := left ; col <= right; col++ {
			result[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			result[index] = matrix[row][right]
			index++
		}

		if left < right && top < bottom {
			for col := right - 1; col > left; col--{
				result[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				result[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return result
}