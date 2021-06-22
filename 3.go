package leetcode

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	wordMap := make(map[byte]int)
	left := 0
	right := 0

	for right < len(s) {
		c := s[right]
		wordMap[c]++
		right++

		for wordMap[c] > 1 {
			d := s[left]
			wordMap[d]--
			left++
		}

		ans = max(ans, right - left)
	}

	return ans
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	win := make(map[byte]int, 0)
	res := 0
	left := 0
	right := 0

	for right < len(s) {
		c := s[right]
		win[c]++
		right++

		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}

		res = max(res, right - left)
	}

	return res
}

