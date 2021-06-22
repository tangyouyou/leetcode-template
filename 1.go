package leetcode

// 两数之和
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, 0)
	numsMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if pos, ok := numsMap[target - nums[i]]; ok {
			result = append(result, pos)
			result = append(result, i)
			break
		} else {
			numsMap[nums[i]] = i
		}
	}

	return result
}
























func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	numsMap := make(map[int]int, 0)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if pos, ok := numsMap[target - nums[i]]; ok {
			result = append(result, pos)
			result = append(result, i)
			break
		} else {
			numsMap[nums[i]] = i
		}
	}

	return result
}
