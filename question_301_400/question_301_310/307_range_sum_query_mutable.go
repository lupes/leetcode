package question_301_310

// 307. 区域和检索 - 数组可修改
// https://leetcode-cn.com/problems/range-sum-query-mutable/
// Topics: 树状数组 线段树

type NumArray2 struct {
	Nums []int
	Sums []int
	Len  int
}

func Constructor2(nums []int) NumArray2 {
	self := NumArray2{
		Nums: nums,
		Len:  len(nums),
	}
	self.Sums = make([]int, self.Len+1)
	self.Sums[0] = 0
	for i, n := range nums {
		self.Sums[i+1] = n + self.Sums[i]
	}
	return self
}

func (this *NumArray2) SumRange(i int, j int) int {
	if i < 0 || j >= this.Len {
		return 0
	}
	return this.Sums[j+1] - this.Sums[i]
}

func (this *NumArray2) Update(i int, val int) {
	diff := val - this.Nums[i]
	this.Nums[i] = val
	for ; i < this.Len; i++ {
		this.Sums[i+1] = this.Sums[i+1] + diff
	}
}
