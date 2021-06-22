package leetcode

import "math"

//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func maxProfit123(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy1 := math.MinInt32
	buy2 := math.MinInt32
	sell1 := 0
	sell2 := 0

	for i := 0; i < len(prices); i++ {
		p := prices[i]
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1 + p)
		buy2 = max(buy2, sell1 - p)
		sell2 = max(sell2, buy2 + p)
	}

	return sell2
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
