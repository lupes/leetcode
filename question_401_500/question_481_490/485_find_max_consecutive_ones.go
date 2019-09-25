package question_481_490

// 485. 最大连续1的个数
// https://leetcode-cn.com/problems/max-consecutive-ones/

func findMaxConsecutiveOnes(nums []int) int {
	var max, t = 0, 0
	for _, n := range nums {
		if n == 1 {
			t++
			if max < t {
				max = t
			}
		} else {
			t = 0
		}
	}
	return max
}
