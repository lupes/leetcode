package question_591_600

// 594. 最长和谐子序列
// https://leetcode-cn.com/problems/longest-harmonious-subsequence/

func findLHS(nums []int) int {
	var tmp = make(map[int]int)
	for _, n := range nums {
		tmp[n]++
	}
	var max = 0
	for _, n := range nums {
		n1 := tmp[n]
		if t, ok := tmp[n-1]; ok && n1+t > max {
			max = n1 + t
		}
		if t, ok := tmp[n-1]; ok && n1+t > max {
			max = n1 + t
		}
	}
	return max
}
