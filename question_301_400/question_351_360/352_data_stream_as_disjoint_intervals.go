package question_351_360

// 352. 将数据流变为多个不相交区间
// https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals
// Topics: 二分查找 设计 有序集合

type SummaryRanges struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor2() SummaryRanges {
	return SummaryRanges{
		nums: []int{-1},
	}
}

func (this *SummaryRanges) AddNum(val int) {
	if val+1 > len(this.nums) {
		tmp := make([]int, val+1)
		copy(tmp, this.nums)
		this.nums = tmp
	}
	this.nums[val] = val
}

func (this *SummaryRanges) GetIntervals() [][]int {
	var res [][]int
	var flag bool
	for i, num := range this.nums {
		if num == i {
			if !flag {
				res = append(res, []int{i, i})
				flag = true
			} else {
				res[len(res)-1][1] = i
			}
		} else {
			flag = false
		}
	}
	return res
}
