package question_1991_2000

// 1992. 找到所有的农场组
// https://leetcode-cn.com/problems/find-all-groups-of-farmland/
// Topic: 深度优先搜索 广度优先搜索 数组 矩阵

func findFarmland(land [][]int) [][]int {
	rl, cl := len(land), len(land[0])
	var res [][]int
	for i, row := range land {
		for j, n := range row {
			if n == 1 {
				tmp := []int{i, j, i, j}
				for a := i; a < rl && land[a][j] == 1; a++ {
					tmp[2] = a
				}
				for b := j; b < cl && land[i][b] == 1; b++ {
					tmp[3] = b
				}
				for a := i; a <= tmp[2]; a++ {
					for b := j; b <= tmp[3]; b++ {
						land[a][b] = 2
					}
				}
				res = append(res, tmp)
			}
		}
	}
	return res
}
