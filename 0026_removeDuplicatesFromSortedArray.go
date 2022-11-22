package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	cur := nums[0]
	point := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			continue
		} else {
			cur = nums[i]
			point++
			nums[point] = cur
		}
	}
	return point + 1
}
