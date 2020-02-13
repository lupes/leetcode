package question_821_830

import (
	"sort"
)

// 825. 适龄的朋友
// https://leetcode-cn.com/problems/friends-of-appropriate-ages
// Topics: 数组

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var res = 0
	for i := len(ages) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if ages[i]/2+7 >= ages[j] {
				break
			}
			res++
			if ages[i] == ages[j] {
				res++
			}
		}
	}
	return res
}
