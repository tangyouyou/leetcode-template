package leetcode

import "sort"

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//输入：candidates = [2,3,6,7], target = 7

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	backTrack(candidates, &res, []int{}, target, 0)

	return res
}

func backTrack(candidates []int, res *[][]int, temp []int, target, start int) {
	if target < 0 {
		return
	}
	//  找到
	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := start; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		backTrack(candidates, res, temp, target - candidates[i], i)
		temp = temp[: len(temp) - 1]
	}
}