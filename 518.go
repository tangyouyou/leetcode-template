func change(amount int, coins []int) int {
	//无限个
	n := len(coins)
	f := make([]int, amount+1)
	f[0] = 1
	//f[i][j] 前i个硬币,组成j的方法数
	for i := 1; i <= n; i++ {
		a := coins[i-1]
		for j := a; j <= amount; j++ {
			f[j] += f[j-a]
		}
	}
	return f[amount]
}