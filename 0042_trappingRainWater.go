package leetcode

func trap(height []int) int {
	maxHeight := 0
	for _, hei := range height {
		if hei > maxHeight {
			maxHeight = hei
		}
	}
	count := 0
	floor := 1
	for {
		if maxHeight < floor {
			break
		}
		val := getFloor(height, floor)
		count += val
		floor++
	}
	return count
}

func getFloor(height []int, floor int) int {
	count := 0
	l := 0
	r := len(height) - 1
	for l < len(height)-1 {
		if height[l] < floor {
			l++
		} else {
			break
		}
	}
	for r >= 0 {
		if height[r] < floor {
			r--
		} else {
			break
		}
	}
	l++
	for l < r {
		if height[l] < floor {
			count++
		}
		l++
	}
	return count
}
