package question_291_300

// 295. 数据流的中位数
// https://leetcode-cn.com/problems/find-median-from-data-stream
// Topics: 堆 设计

type MedianFinder struct {
	nums []int
	len  int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0, 100),
		len:  0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
	this.len++
	for i := this.len - 1; i > 0 && this.nums[i] < this.nums[i-1]; i-- {
		this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.len == 0 {
		return 0
	} else if this.len%2 == 0 {
		return float64(this.nums[this.len/2-1]+this.nums[this.len/2]) / 2
	}
	return float64(this.nums[this.len/2])
}
