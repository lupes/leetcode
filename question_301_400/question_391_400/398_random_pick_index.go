package question_391_400

import (
	"math/rand"
)

// 398. 随机数索引
// https://leetcode-cn.com/problems/random-pick-index/
// Topics: 蓄水池抽样

type Solution struct {
	indexs map[int][]int
	nums   map[int]int
}

func Constructor(nums []int) Solution {
	var indexs = make(map[int][]int)
	var ns = make(map[int]int)
	for i, n := range nums {
		indexs[n] = append(indexs[n], i)
		ns[n]++
	}
	return Solution{
		indexs: indexs,
		nums:   ns,
	}
}

func (this *Solution) Pick(target int) int {
	return this.indexs[target][rand.Intn(this.nums[target])]
}
