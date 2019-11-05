package question_591_600

// 598. 范围求和 II
// https://leetcode-cn.com/problems/range-addition-ii/

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		if op[0] < m {
			m = op[0]
		}
		if op[1] < n {
			n = op[1]
		}
	}
	return m * n
}
