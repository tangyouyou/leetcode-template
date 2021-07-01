// 最长子数组的长度
// 暴力解法
func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	ans := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			k := 0
			for i + k < len(A) && j + k < len(B) && A[i+k] == B[j+k] {
				k++
			}
			if k > ans {
				ans = k
			}
		}
	}

	return ans
}


// 思路，动态规划 dp[i][j] = dp[i+1][j+1] + 1
// 从尾向前进行遍历,因为需要先计算dp[i+1][j+1]
func findLength2(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	lenA := len(A)
	lenB := len(B)
	dp := make([][]int, lenA+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lenB+1)
	}
	ans := 0
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}

	return ans
}