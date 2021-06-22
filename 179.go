package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
func largestNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)

	for i := 0; i < n; i++ {
		strs[i] = fmt.Sprintf("%d", nums[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i] + strs[j] > strs[j] + strs[i]
	})

	if strs[0][0] == '0' {
		return "0"
	}

	return strings.Join(strs, "")
}
