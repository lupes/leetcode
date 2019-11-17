package question_301_310

// 303. 区域和检索 - 数组不可变
// https://leetcode-cn.com/problems/range-sum-query-immutable/
// Topics: 动态规划

type NumArray struct {
	Sums []int
	Len  int
}

func Constructor(nums []int) NumArray {
	a := NumArray{
		Len: len(nums),
	}
	a.Sums = make([]int, a.Len+1)
	a.Sums[0] = 0
	for i, n := range nums {
		a.Sums[i+1] = n + a.Sums[i]
	}
	return a
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j >= this.Len {
		return 0
	}
	return this.Sums[j+1] - this.Sums[i]
}
