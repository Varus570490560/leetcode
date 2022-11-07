package leetcode

func numberOfDays(year int, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		} else {
			return 28
		}
	default:
		return 0
	}
}

func isLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}
