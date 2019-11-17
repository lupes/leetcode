package question_491_500

// 498. 对角线遍历
// https://leetcode-cn.com/problems/diagonal-traverse/
// Topics:

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var res []int
	r, c, rl, cl, flag := 0, 0, len(matrix), len(matrix[0]), true
	for r < rl && c < cl {
		res = append(res, matrix[r][c])
		if flag {
			if r-1 >= 0 && c+1 < cl {
				r -= 1
				c += 1
			} else if r-1 < 0 && c+1 < cl {
				flag = false
				c += 1
			} else if c+1 >= cl {
				flag = false
				r += 1
			}
		} else {
			if r+1 < rl && c-1 >= 0 {
				r += 1
				c -= 1
			} else if r+1 < rl && c-1 < 0 {
				flag = true
				r += 1
			} else if r+1 >= rl {
				flag = true
				c += 1
			}
		}
	}
	return res
}
