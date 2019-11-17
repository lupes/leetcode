package question_221_230

// 229. 求众数 II
// https://leetcode-cn.com/problems/majority-element-ii/

func majorityElement(nums []int) []int {
	l := len(nums) / 3
	m := make(map[int]int)
	var res []int
	for _, n := range nums {
		m[n]++
	}
	for k, n := range m {
		if n > l {
			res = append(res, k)
		}
	}
	return res
}
