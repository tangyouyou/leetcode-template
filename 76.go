package leetcode

import "math"

//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"

func minWindow(s string, t string) string {
	res := ""
	cnt := math.MaxInt32
	l := 0
	r := 0
	wordMap := make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		wordMap[t[i]]++
	}

	for r < len(s) {
		wordMap[s[r]]--
		for check(wordMap) {
			if r-l+1 < cnt {
				cnt = r - l + 1
				res = s[l : r+1]
			}
			wordMap[s[l]]++
			l++
		}
		r++
	}
	return res
}

func check(wordMap map[byte]int) bool {
	for _, v := range wordMap {
		if v > 0 {
			return false
		}
	}
	return true
}

//func minWindow(s string, t string) string {
//	var res string
//	cnt := math.MaxInt32
//	hashMap := make(map[byte]int)
//	l := 0
//	r := 0
//	for i := 0; i < len(t); i++ {
//		hashMap[t[i]]++
//	}
//	for r < len(s) {
//		hashMap[s[r]]--
//		for check(hashMap) {
//			if r-l+1 < cnt {
//				cnt = r - l + 1
//				res = s[l : r+1]
//			}
//			hashMap[s[l]]++
//			l++
//		}
//		r++
//	}
//	return res
//}
//
//func check(hashMap map[byte]int) bool {
//	for _, v := range hashMap {
//		if v > 0 {
//			return false
//		}
//	}
//	return true
//}
