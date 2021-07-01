//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果
//一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。 
//
// [["a","b","c","e"], 
//["s","f","c","s"], 
//["a","d","e","e"]] 
//
// 但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。 
//
// 
//
// 示例 1： 
//
// 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "AB
//CCED"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) == true {
				return true
			}
		}
	}

	return false
}


func dfs(board [][]byte, word string, x int, y int, pos int) bool {
	if board[x][y] != word[pos] {
		return false
	}

	if pos == len(word) - 1 {
		return true
	}

	temp := board[x][y]
	board[x][y] = byte(' ')
	if x - 1 >= 0 {
		if dfs(board, word, x-1, y, pos + 1) {
			return true
		}
	}
	if x + 1 < len(board) {
		if dfs(board, word, x+1, y, pos + 1) {
			return true
		}
	}
	if y - 1 >= 0 {
		if dfs(board, word, x, y-1, pos + 1) {
			return true
		}
	}
	if y + 1 < len(board[0]) {
		if dfs(board, word, x, y+1, pos + 1) {
			return true
		}
	}
	board[x][y] = temp

	return false
}