package leetcode

func subtractProductAndSum(n int) int {
	numBitsArray := make([]int, 0)
	for n != 0 {
		numBitsArray = append(numBitsArray, n%10)
		n /= 10
	}
	sum := 0
	mut := 1
	for _, bit := range numBitsArray {
		sum += bit
		mut *= bit
	}
	return mut - sum
}
