package common

func countingSort(nums []int) []int {
	var flag, res = make([]int, 10001), make([]int, len(nums))
	for _, n := range nums {
		flag[n]++
	}
	var sum = 0
	for i, n := range flag {
		sum += n
		flag[i] = sum
	}
	for _, n := range nums {
		flag[n]--
		res[flag[n]] = n
	}
	return res
}
