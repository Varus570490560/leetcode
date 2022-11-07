package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	maxLength := 1
	curLength := 1
	pointA := 0
	pointB := 0
	set := make(map[byte]bool, 0)
	set[s[pointA]] = true
	for {
		pointB++
		if pointB == len(s) {
			return maxLength
		} else if set[s[pointB]] == false {
			curLength++
			if maxLength < curLength {
				maxLength = curLength
			}
			set[s[pointB]] = true
		} else {
			pointB--
			set[s[pointA]] = false
			pointA++
			curLength--
		}
	}
}
