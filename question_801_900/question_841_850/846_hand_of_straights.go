package question_841_850

import (
	"sort"
)

// 846. 一手顺子
// https://leetcode-cn.com/problems/hand-of-straights
// Topics: None

func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}

	sort.Ints(hand)

	var flag = make(map[int][]int)
	for _, n := range hand {
		arr := flag[n]
		if len(arr) == 0 {
			flag[n+1] = append(flag[n+1], 1)
		} else {
			t := arr[0] + 1
			flag[n] = arr[1:]
			if t < W {
				flag[n+1] = append(flag[n+1], t)
			}
			if len(flag[n]) == 0 {
				delete(flag, n)
			}
		}
	}
	return len(flag) == 0
}
