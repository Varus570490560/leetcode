package leetcode

func largestAltitude(gain []int) int {
	var res, cur = 0, 0
	for _, val := range gain {
		cur += val
		if cur > res {
			res = cur
		}
	}
	return res
}
