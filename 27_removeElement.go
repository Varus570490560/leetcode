package leetcode

func removeElement(nums []int, val int) int {
	var offset int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			offset++
		} else {
			nums[i-offset] = nums[i]
		}
	}
	return len(nums) - offset
}
