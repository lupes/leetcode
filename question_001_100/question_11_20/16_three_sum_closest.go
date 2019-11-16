package question_11_20

import (
	"math"
	"sort"
)

// 16. 最接近的三数之和
// https://leetcode-cn.com/problems/3sum-closest

func threeSumClosest(nums []int, target int) int {
	size := len(nums)
	if size < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < size-2; i++ {
		l, r := i+1, size-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(tmp-target)) < math.Abs(float64(res-target)) {
				res = tmp
			}
			if tmp > target {
				r--
			} else if tmp < target {
				l++
			} else {
				return target
			}
		}
	}
	return res
}
