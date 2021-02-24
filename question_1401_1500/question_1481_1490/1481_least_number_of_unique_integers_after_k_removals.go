package question_1481_1490

import (
	"sort"
)

// 1481. 不同整数的最少数目
// https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals/
// Topics: 排序 数组

func findLeastNumOfUniqueInts(arr []int, k int) int {
	var flag = make(map[int]int, 0)
	for _, n := range arr {
		flag[n]++
	}

	var tmp = make([][2]int, 0, len(flag))
	for k, v := range flag {
		tmp = append(tmp, [2]int{k, v})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] < tmp[j][1]
	})

	for k > 0 && len(tmp) > 0 {
		if tmp[0][1] > k {
			tmp[0][1] -= k
			k = 0
		} else {
			k -= tmp[0][1]
			tmp = tmp[1:]
		}
	}

	return len(tmp)
}
