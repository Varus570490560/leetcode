package leetcode

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int, 0)
	for idx, value := range nums {
		anotherIdx, ok := hashTable[target-value]
		if !ok {
			hashTable[value] = idx
		} else {
			return []int{
				idx,
				anotherIdx,
			}
		}
	}
	return []int{
		-1,
		-1,
	}
}
