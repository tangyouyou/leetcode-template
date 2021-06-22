package leetcode

import "sort"

// 三数之和

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]

		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			n2 := nums[left]
			n3 := nums[right]

			if n1+n2+n3 == 0 {
				result = append(result, []int{n1, n2, n3})
				for left < right && n2 == nums[left] {
					left++
				}
				for left < right && n3 == nums[right] {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}


























func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			n2 := nums[left]
			n3 := nums[right]

			if n1 + n2 + n3 == 0 {
				result = append(result, []int{n1, n2, n3})

				for left < right && n2 == nums[left] {
					left++
				}
				for left < right && n3 == nums[right] {
					right--
				}
			} else if n1 + n2 + n3 < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}