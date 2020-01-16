package question_701_710

import (
	"sort"
)

// 703. 数据流中的第K大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
// Topics: 堆

type KthLargest struct {
	k    int
	nums []int
}

func Constructor2(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	for i, n := range this.nums {
		if n > val {
			tmp := make([]int, 0, len(this.nums)+1)
			tmp = append(tmp, this.nums[:i]...)
			tmp = append(tmp, val)
			tmp = append(tmp, this.nums[i:]...)
			this.nums = tmp
			goto Next
		}
	}
	this.nums = append(this.nums, val)
Next:
	return this.nums[len(this.nums)-this.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
