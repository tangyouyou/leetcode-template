func demo_200(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs_200(grid, i, j)
			}
		}
	}

	return res
}

func dfs_200(grid [][]int, x int, y int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		dfs_200(grid, x-1, y)
		dfs_200(grid, x+1, y)
		dfs_200(grid, x, y-1)
		dfs_200(grid, x, y+1)
	}
}


