package question_1411_1420

// 1413. 逐步求和得到正数的最小值
// https://leetcode-cn.com/problems/minimum-value-to-get-positive-step-by-step-sum/
// Topics: 数组

func minStartValue(nums []int) int {
	var min, sum = 1, 0
	for _, n := range nums {
		sum += n
		if min+sum < 1 {
			min = 1 - sum
		}
	}
	return min
}
