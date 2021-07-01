func demo_42(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(right) - 1
	leftMax := 0
	rightMax := 0
	ans := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans = ans + leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans = ans + rightMax - height[right]
			}
			right--
		}
	}

	return ans
}