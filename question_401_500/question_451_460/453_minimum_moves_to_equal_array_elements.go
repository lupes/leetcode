package question_451_460

// 453. 最小移动次数使数组元素相等
// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements
// Topics: 数学

func minMoves(nums []int) int {
	var min, res = nums[0], 0
	for i, n := range nums {
		if n < min {
			res += (min - n) * i
			min = n
		} else {
			res += n - min
		}
	}
	return res
}
