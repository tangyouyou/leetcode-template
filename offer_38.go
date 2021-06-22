func permutation(s string) []string {
if len(s) == 0 {
return []string{}
}
c := []byte(s)
result := make([]string, 0)
wordMap := make(map[string]bool, 0)
dfs38(c, 0, wordMap)
for k := range wordMap {
result = append(result, k)
}

return result
}

func dfs38(c []byte, i int, wordMap map[string]bool) {
	if i == len(c)-1 {
		wordMap[string(c)] = true
	return
	} else {
	for j := i; j < len(c); j++ {
		c[i], c[j] = c[j], c[i]
		dfs38(c, i+1, wordMap)
		c[i], c[j] = c[j], c[i]
	}
	}
}