func demo_122(profile []int) int {
	if len(profile) == 0 {
		return 0
	}
	iMax := 0

	for i := 1; i < len(profile); i++ {
		if profile[i] > profile[i-1] {
			iMax += profile[i] - profile[i-1]
		}
	}

	return iMax
}