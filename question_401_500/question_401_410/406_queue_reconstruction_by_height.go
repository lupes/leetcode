package question_401_410

import (
	"sort"
)

// 406. 根据身高重建队列
// https://leetcode-cn.com/problems/queue-reconstruction-by-height
// Topics: 贪心算法

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	for i, p := range people {
		if p[1] != i {
			for j := i; j > p[1]; j-- {
				people[j-1], people[j] = people[j], people[j-1]
			}
		}
	}
	return people
}
