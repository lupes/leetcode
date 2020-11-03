package question_1331_1340

// 1351. 统计有序矩阵中的负数
// https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended/
// Topics: 数组 二分查找

func countNegatives(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	var res = 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] < 0 {
				res += c - j
				break
			}
		}
	}
	return res
}
