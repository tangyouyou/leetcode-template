package leetcode


func nextPermutation(nums []int)  {
	// 1. 全降序，代表是最大排列，需要转换成全升序，代表最小排列

	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2
	j := len(nums) - 1
	k := len(nums) - 1

	// 从后向前遍历，找到到第一个相对升序, nums[i] >= nums[j], 代表没有找到，需要继续往前找
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 从 j 到 尾部，查找第一个 a[k] > a[i]
	if i >= 0 {
		// find: A[i] < A[k]
		for nums[i] >= nums[k] {
			k--
		}
		nums[i],nums[k] = nums[k],nums[i]
	}

	// 从 j 到尾部，从降序调整为升序
	i = j
	j = len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}
