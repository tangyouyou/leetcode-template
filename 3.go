func demo_3(s string) int {
	if len(s) == 0 {
		return 0
	}

	win := make(map[byte]int)
	ans := 0
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

		if right - left > ans {
			ans = right - left
		}
	}

	return ans
}	