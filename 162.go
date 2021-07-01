// 峰值元素是指其值大于左右相邻值的元素。

// 给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

// 你可以假设 nums[-1] = nums[n] = -∞ 。

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	pos := 0

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] < nums[i+1] {
			pos++
			continue
		} else {
			break
		}
	}

	return pos
}
