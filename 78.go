// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)

	backTrack(nums, 0, list, &result)

	return result
}

func backTrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++{
		list = append(list, nums[i])
		backTrack(nums, i+1, list, result)
		list = list[0:len(list)-1]
	}
}