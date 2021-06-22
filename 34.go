package leetcode

func searchRange(nums []int, target int) []int {
	left := binary(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := binary(nums, target+1) - 1

	return []int{left, right}
}

func binary(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else  {
			right = mid - 1
		}
	}

	return left
}