package leetcode

func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		if !getNBit(strs, i) {
			return strs[0][0:i]
		}
	}
}
func getNBit(strs []string, n int) bool {
	if len(strs[0]) <= n {
		return false
	}
	c := strs[0][n]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= n {
			return false
		}
		if strs[i][n] != c {
			return false
		}
	}
	return true
}
