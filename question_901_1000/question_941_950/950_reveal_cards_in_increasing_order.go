package question_941_950

import (
	"sort"
)

// 950. 按递增顺序显示卡牌
// https://leetcode-cn.com/problems/reveal-cards-in-increasing-order
// Topics: 数组

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	var flag = make([]int, len(deck), len(deck)*2)
	for i := 0; i < len(deck); i++ {
		flag[i] = i
	}
	var res = make([]int, len(deck))
	for i := 0; len(flag) > 0; i++ {
		res[flag[0]] = deck[i]

		flag = flag[1:]
		if len(flag) > 1 {
			flag = append(flag, flag[0])
			flag = flag[1:]
		}
	}
	return res
}
