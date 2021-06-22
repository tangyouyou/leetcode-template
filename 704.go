package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

