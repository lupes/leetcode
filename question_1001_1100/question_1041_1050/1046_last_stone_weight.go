package question_1041_1050

import (
	"sort"
)

// 1046. 最后一块石头的重量
// https://leetcode-cn.com/problems/last-stone-weight
// Topics: 堆 贪心算法

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
		if stones[0] == stones[1] {
			stones = stones[2:]
		} else {
			stones[1] = stones[0] - stones[1]
			stones = stones[1:]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
