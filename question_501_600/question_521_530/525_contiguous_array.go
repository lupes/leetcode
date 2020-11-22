package question_521_530

// 525. 连续数组
// https://leetcode-cn.com/problems/contiguous-array
// Topics: 哈希表

func findMaxLength(nums []int) int {
	var flag, zeroSum, oneSum, max, a, b = make(map[int]int), 0, 0, 0, 0, 0
	for i, n := range nums {
		if n == 0 {
			zeroSum++
		} else {
			oneSum++
		}
		a, b = oneSum-zeroSum, zeroSum-oneSum
		if a == 0 && i > max {
			max = i + 1
		}
		if j, ok := flag[b]; ok && i-j > max {
			max = i - j
		}
		if _, ok := flag[b]; !ok {
			flag[b] = i
		}
	}
	return max
}
