package question_1561_1570

// 1561. 你可以获得的最大硬币数目
// https://leetcode-cn.com/problems/maximum-number-of-coins-you-can-get/
// Topics: 排序

import (
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	var res = 0
	for i := len(piles) / 3; i < len(piles); i += 2 {
		res += piles[i]
	}
	return res
}
