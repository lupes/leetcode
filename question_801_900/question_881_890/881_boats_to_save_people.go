package question_881_890

import (
	"sort"
)

// 881. 救生艇
// https://leetcode-cn.com/problems/boats-to-save-people
// Topics: 贪心算法 双指针

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var res = 0
	for left, right := 0, len(people)-1; right >= left; {
		if people[left]+people[right] <= limit {
			left++
		}
		right, res = right-1, res+1
	}
	return res
}
