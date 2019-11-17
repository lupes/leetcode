package question_521_530

import (
	"math/rand"
)

// 528. 按权重随机选择
// https://leetcode-cn.com/problems/random-pick-with-weight/
// Topics: 二分查找 None

type Solution struct {
	Sums []int
	Sum  int
	Len  int
}

func Constructor(w []int) Solution {
	var l = len(w)
	var sums = make([]int, l+1)
	var sum = 0
	for i, n := range w {
		sums[i+1] = sums[i] + n
		sum += n
	}
	return Solution{
		Sums: sums,
		Sum:  sum,
		Len:  l,
	}
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.Sum) + 1
	for l, r, c := 0, this.Len, this.Len/2; l < r; c = (l + r) / 2 {
		if this.Sums[c] < n && n <= this.Sums[c+1] {
			return c
		}
		if this.Sums[c] > n {
			r = c
		} else if this.Sums[c] < n {
			l = c
		} else {
			return c - 1
		}
	}
	return -1
}
