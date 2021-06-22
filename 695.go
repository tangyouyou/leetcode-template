package leetcode

//给定一个包含了一些 0 和 1 的非空二维数组 grid 。
//一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	iMax := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				tmp := dfs695(grid, i, j)
				if tmp > iMax {
					iMax = tmp
				}
			}
		}
	}

	return iMax
}

func dfs695(grid [][]int, x, y int) int {
	if grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	res := 1

	if x - 1 >= 0 {
		res = res + dfs695(grid, x - 1, y)
	}
	if x + 1 < len(grid) {
		res = res + dfs695(grid, x + 1, y)
	}

	if y - 1 >= 0 {
		res = res + dfs695(grid, x, y - 1)
	}
	if y + 1 < len(grid[0]) {
		res = res + dfs695(grid, x, y + 1)
	}

	return res
}

