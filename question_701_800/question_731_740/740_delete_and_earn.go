package question_731_740

import (
	"sort"
)

// 740. 删除与获得点数
// https://leetcode-cn.com/problems/delete-and-earn
// Topics: 动态规划

func deleteAndEarn(nums []int) int {
	var tmp = make(map[int]int)
	var nums2 = []int{0, 0}
	for _, n := range nums {
		if _, ok := tmp[n]; !ok {
			nums2 = append(nums2, n)
		}
		tmp[n]++
	}
	sort.Ints(nums2)
	var flag = make([]int, len(nums2))
	for i := 2; i < len(nums2); i++ {
		max, l := 0, i
		if nums2[i] == nums2[i-1]+1 {
			l--
		}
		for j := 2; j < l; j++ {
			if flag[j] > max {
				max = flag[j]
			}
		}
		flag[i] = nums2[i]*tmp[nums2[i]] + max
	}
	if flag[len(flag)-1] > flag[len(flag)-2] {
		return flag[len(flag)-1]
	} else {
		return flag[len(flag)-2]
	}
}
