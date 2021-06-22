package leetcode
//
//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}