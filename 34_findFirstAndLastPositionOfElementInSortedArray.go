package leetcode

func searchRange(nums []int, target int) []int {
	l := 0
	idx := -1
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			idx = m
			break
		}
	}
	if idx == -1 {
		return []int{
			-1,
			-1,
		}
	}
	l = idx
	r = idx
	for l-1 >= 0 && nums[l-1] == target {
		l--
	}
	for r+1 < len(nums) && nums[r+1] == target {
		r++
	}
	return []int{
		l,
		r,
	}
}
