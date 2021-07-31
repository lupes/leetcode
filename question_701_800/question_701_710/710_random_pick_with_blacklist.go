package question_701_710

import (
	"math/rand"
	"sort"
)

// 710. 黑名单中的随机数
// https://leetcode-cn.com/problems/random-pick-with-blacklist
// Topics: 排序 哈希表 二分查找 None

type Solution struct {
	max int
	m   map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	sort.Ints(blacklist)
	var tmp = make(map[int]int, len(blacklist))
	for _, n := range blacklist {
		tmp[n] = -1
	}
	max := N - len(blacklist)
	for i, j := max, 0; i < N; i++ {
		if _, ok := tmp[i]; !ok {
			tmp[blacklist[j]], j = i, j+1
		}
	}
	return Solution{
		max: max,
		m:   tmp,
	}
}

func (this *Solution) Pick() int {
	n := rand.Intn(this.max)
	if t, ok := this.m[n]; ok {
		n = t
	}
	return n
}
