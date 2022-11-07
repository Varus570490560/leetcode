package leetcode

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	} else if x < 0 {
		return false
	} else {
		var xCopy int = x
		var reverse int = 0
		for xCopy != 0 {
			reverse *= 10
			reverse += xCopy % 10
			xCopy /= 10
		}
		if reverse == x {
			return true
		}
		return false

	}
}
