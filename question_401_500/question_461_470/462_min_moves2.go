package question_461_470

// 462. 最少移动次数使数组元素相等 II
// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/

func minMoves2(nums []int) int {
	var l, r = len(nums), 0
	for i := 0; i < l-1; i++ {
		m, t := nums[0], 0
		for j := 0; j < l-i; j++ {
			if nums[j] > m {
				t = j
				m = nums[j]
			}
		}
		nums[t], nums[l-i-1] = nums[l-i-1], nums[t]
	}
	for i, j := 0, l-1; j > i; i, j = i+1, j-1 {
		r += nums[j] - nums[i]
	}
	return r
}
