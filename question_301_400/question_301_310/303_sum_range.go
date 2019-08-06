package question_301_310

// 303. 区域和检索 - 数组不可变
// https://leetcode-cn.com/problems/range-sum-query-immutable/

type NumArray struct {
	Nums []int
	Len  int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Nums: nums,
		Len:  len(nums),
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || j >= this.Len {
		return 0
	}
	res := 0
	for _, n := range this.Nums[i : j+1] {
		res += n
	}
	return res
}
