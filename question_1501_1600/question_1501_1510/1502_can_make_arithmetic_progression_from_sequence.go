package question_1501_1510

// 1502. 判断能否形成等差数列
// https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/
// Topics: 排序 数组

import (
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}
