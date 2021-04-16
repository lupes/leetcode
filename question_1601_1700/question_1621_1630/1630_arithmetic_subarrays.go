package question_1621_1630

// 1630. 等差子数组
// https://leetcode-cn.com/problems/arithmetic-subarrays/
// Topics: 排序

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var res = make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		res[i] = true
		tmp := make([]int, r[i]-l[i]+1)
		copy(tmp, nums[l[i]:r[i]+1])
		if len(tmp) == 1 {
			continue
		}
		sort.Ints(tmp)
		d := tmp[1] - tmp[0]
		for j := 2; j < len(tmp); j++ {
			if tmp[j]-tmp[j-1] != d {
				res[i] = false
				break
			}
		}
	}
	return res
}
