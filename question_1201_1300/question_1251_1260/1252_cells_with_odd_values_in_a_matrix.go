package question_1251_1260

// 1252. 奇数值单元格的数目
// https://leetcode-cn.com/problems/cells-with-odd-values-in-a-matrix
// Topics: 数组

func oddCells(n int, m int, indices [][]int) int {
	var flag = make([][]int, n)
	for i := range flag {
		flag[i] = make([]int, m)
	}
	for _, ind := range indices {
		for i := 0; i < m; i++ {
			flag[ind[0]][i]++
		}
		for i := 0; i < n; i++ {
			flag[i][ind[1]]++
		}
	}
	var res = 0
	for _, row := range flag {
		for _, c := range row {
			if c%2 == 1 {
				res++
			}
		}
	}
	return res
}
