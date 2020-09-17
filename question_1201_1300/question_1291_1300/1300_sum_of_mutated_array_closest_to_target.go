package question_1291_1300

import (
	"sort"
)

// 1300. 转变数组后最接近目标值的数组和
// https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/
// Topics: 二分查找 数组

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	l, pre := len(arr), 0
	for i, n := range arr {
		cur := pre + (l-i)*n
		if cur >= target {
			v := (target - pre) / (l - i)
			low := pre + v*(l-i)
			high := low + l - i
			if target-low <= high-target {
				return v
			} else {
				return v + 1
			}
		}
		pre += n
	}
	return arr[l-1]
}
