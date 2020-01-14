package question_701_710

// 703. 数据流中的第K大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
// Topics: 堆

type KthLargest struct {
}

func Constructor2(k int, nums []int) KthLargest {
	return KthLargest{}
}

func (this *KthLargest) Add(val int) int {
	return 0
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
