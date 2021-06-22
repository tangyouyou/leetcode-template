package leetcode

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row := len(grid) - 1
	cols := len(grid[0]) - 1
	res := 0
	for i := 0; i <= row; i ++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				res++
				dfs200(grid, i, j)
			}
		}
	}

	return res
}

func dfs200(grid [][]byte, x int, y int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		dfs200(grid, x+1, y)
		dfs200(grid, x-1, y)
		dfs200(grid, x, y+1)
		dfs200(grid, x, y-1)
	}
}