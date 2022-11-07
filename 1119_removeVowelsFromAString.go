package leetcode

func removeVowels(s string) string {
	res := ""
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			continue
		} else {
			res = res + string(char)
		}
	}
	return res
}
