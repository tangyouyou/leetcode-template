// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其和 ≥ target 的长度最小的连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
// 如果不存在符合条件的子数组，返回 0 。

// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt32
	n := len(nums)

	for i := 0; i < n; i++ {
		sum := 0	
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target {
				ans := min(ans, j - i + 1)
				break
			}
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}


// 滑动窗口法
func minSubArrayLen(s int, nums []int) int {
    if len(nums) == 0 {
    	return 0
    }
    n := len(nums)

    left := 0
    right := 0
    ans := math.MaxInt32
    sum := 0

    for right < n {
    	sum += nums[right]
    	if sum >= s {
    		sum -= nums[left]
    		ans = min(ans, right - left + 1)
    		left++
    	}
    	right++
    }

    if ans == math.MaxInt32 {
    	return 0
    }

    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
