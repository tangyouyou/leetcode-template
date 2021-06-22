package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	data := map[byte]byte {
		'(' : ')',
		'{' : '}',
		'[' : ']',
	}

	stack := make([]byte, 0, len(s) / 2)

	for i := 0; i < len(s); i++ {
		q := s[i]
		// 判断是否有效的符号
		if _, ok := data[q]; ok {
			stack = append(stack, q)
		} else {
			// 匹配到右括号
			if len(stack) > 0 && (data[stack[len(stack) - 1]] == q) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false

}