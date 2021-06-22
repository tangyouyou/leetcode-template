package leetcode

import "strconv"

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1

	addon := 0
	res := ""
	n1 := 0
	n2 := 0
	tmp := 0

	for i >= 0 || j >= 0 {
		if i >= 0 {
			n1, _ = strconv.Atoi(num1[i:i+1])
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2, _ = strconv.Atoi(num2[j:j+1])
		} else {
			n2 = 0
		}

		tmp = n1 + n2 + addon
		addon = tmp / 10

		res = strconv.Itoa(tmp % 10) + res

		i--
		j--
	}

	if addon == 1 {
		res = "1" + res
	}

	return res
}
