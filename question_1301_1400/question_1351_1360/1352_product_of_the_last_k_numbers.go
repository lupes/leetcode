package question_1331_1340

// 1352. 最后 K 个数的乘积
// https://leetcode-cn.com/problems/product-of-the-last-k-numbers/
// Topics: 设计 数组

type ProductOfNumbers struct {
	l        int
	lastZero int
	nums     []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		l:    0,
		nums: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.lastZero = 0
		this.nums = append(this.nums, 1)
	} else {
		this.lastZero++
		this.nums = append(this.nums, this.nums[this.l]*num)
	}
	this.l++
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > this.lastZero {
		return 0
	}
	return this.nums[this.l] / this.nums[this.l-k]
}
