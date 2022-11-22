package leetcode

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int, 0)
	for i := lowLimit; i <= highLimit; i++ {
		m[getBoxNum(i)]++
	}
	var res int
	for key, _ := range m {
		if res == 0 {
			res = key
		} else {
			if m[res] < m[key] {
				res = key
			}
		}
	}
	return m[res]
}
func getBoxNum(ballNum int) int {
	count := 0
	for ballNum != 0 {
		count += ballNum % 10
		ballNum /= 10
	}
	return count
}
