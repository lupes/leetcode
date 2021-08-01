package question_1311_1320

// 1337. 矩阵中战斗力最弱的 K 行
// https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
// Topics: 数组 二分查找 矩阵 排序 堆

func kWeakestRows(mat [][]int, k int) []int {
	var res, rl, cl = make([]int, 0, k), len(mat), len(mat[0])
	for i := 0; i < cl; i++ {
		for j := 0; j < rl; j++ {
			if mat[j][i] == 0 && (i == 0 || (i > 0 && mat[j][i-1] == 1)) {
				res = append(res, j)
				if len(res) == k {
					return res
				}
			}
		}
	}
	for i := range mat {
		if mat[i][cl-1] == 1 {
			res = append(res, i)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
