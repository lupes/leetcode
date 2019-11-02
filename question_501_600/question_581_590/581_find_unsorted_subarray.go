package question_581_590

import (
	"sort"
)

// 581. 最短无序连续子数组
// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
	l, fl, r, fr := 0, false, len(nums)-1, false
	var tmp = make([]int, r+1)
	copy(tmp, nums)
	sort.Ints(tmp)
	for l <= r && (!fl || !fr) {
		if nums[l] == tmp[l] {
			l++
		} else {
			fl = true
		}
		if nums[r] == tmp[r] {
			r--
		} else {
			fr = true
		}
	}
	if r <= l {
		return 0
	}
	return r - l + 1
}
