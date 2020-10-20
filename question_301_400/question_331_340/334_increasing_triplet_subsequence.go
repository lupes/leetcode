package question_331_340

import (
	"math"
)

// 334. 递增的三元子序列
// https://leetcode-cn.com/problems/increasing-triplet-subsequence
// Topics:

func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n <= a {
			a = n
		} else if n <= b {
			b = n
		} else {
			return true
		}
	}
	return false
}
