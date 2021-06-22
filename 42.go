package leetcode

// 接雨水
func trap(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result = result + leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result = result + rightMax - height[right]
			}
			right--
		}
	}

	return result
}










































func trap(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	result := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result = result + leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result = result + rightMax - height[right]
			}
			right--
		}
	}

	return result
}
