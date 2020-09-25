package question_381_390

import (
	"math/rand"
)

// 384. 打乱数组
// https://leetcode-cn.com/problems/shuffle-an-array
// Topics:

type Solution struct {
	src []int
}

func Constructor(nums []int) Solution {
	return Solution{
		src: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.src
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	var res = make([]int, 0, len(this.src))
	var tmp = make([]int, len(this.src))
	copy(tmp, this.src)
	for len(tmp) > 0 {
		n := rand.Intn(len(tmp))
		res = append(res, tmp[n])
		tmp = append(tmp[:n], tmp[n+1:]...)
	}
	return res
}
