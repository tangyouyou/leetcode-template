package leetcode

//提供一种很容易想到的思路
//核心思想还是缺失值不超过len或者是len+1
//申请一个与原数组等长的数组，遍历然后记下每个在范围内（0~len）出现的数的次数，然后审查该数组中
//是否含有值为0即可。
func firstMissingPositive(nums []int) int {
	maxLen := len(nums)

	if maxLen == 0 {
		return 1
	}

	if maxLen == 1 && nums[0] == 1 {
		return 2
	}

	tmpArr := make([]int, maxLen + 1)

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > maxLen {
			continue
		}
		tmpArr[nums[i]]++
	}

	var j int
	for j = 1; j < len(tmpArr); j++ {
		if tmpArr[j] == 0 {
			return j
		}
	}
	return j
}


