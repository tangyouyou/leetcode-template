// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	valid := true


	for _, values := range prerequisites {
		edges[values[1]] = append(edges[values[1]], values[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs207(i, visited, edges, &valid)
		}
	}

	return valid
}	

func dfs207(u int, visited []int, edges [][]int, valid *bool) {
	visited[u] = 1
	for _, v := range edges[u] {
		if visited[v] == 0 {
			dfs207(v, visited, edges, valid)
			if *valid == false {
				return
			}
		} else if visited[v] == 1 {
			*valid = false
			return
		}
	}
	visited[u] = 2
}