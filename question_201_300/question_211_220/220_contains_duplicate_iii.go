package question_211_220

import "math"

// 220. 存在重复元素 III
// https://leetcode-cn.com/problems/contains-duplicate-iii/
// Topics: 排序 None

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, n := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if math.Abs(float64(n-nums[j])) <= float64(t) {
				return true
			}
		}
	}
	return false
}
