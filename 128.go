package leetcode

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numsMap := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = true
	}
	ans := 0

	for num := range numsMap {
		if numsMap[num-1] == false {
			cur := 1
			tmpNum := num

			for numsMap[tmpNum+1] {
				cur++
				tmpNum++
			}

			ans = max(cur, ans)
		}
	}

	return ans
}
