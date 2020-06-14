package question_1221_1230

// 1222. 可以攻击国王的皇后
// https://leetcode-cn.com/problems/queens-that-can-attack-the-king
// Topics: 数组

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var flag, res = make(map[int][]int), make([][]int, 0)
	for _, queen := range queens {
		flag[queen[0]*10+queen[1]] = queen
	}
	for i := king[0] - 1; i >= 0; i-- {
		if v, ok := flag[i*10+king[1]]; ok {
			res = append(res, v)
			break
		}
	}
	for i := king[0] + 1; i < 8; i++ {
		if v, ok := flag[i*10+king[1]]; ok {
			res = append(res, v)
			break
		}
	}

	for i := king[1] - 1; i >= 0; i-- {
		if v, ok := flag[king[0]*10+i]; ok {
			res = append(res, v)
			break
		}
	}
	for i := king[1] + 1; i < 8; i++ {
		if v, ok := flag[king[0]*10+i]; ok {
			res = append(res, v)
			break
		}
	}

	for i, j := king[0]-1, king[1]-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if v, ok := flag[i*10+j]; ok {
			res = append(res, v)
			break
		}
	}
	for i, j := king[0]+1, king[1]+1; i < 8 && j < 8; i, j = i+1, j+1 {
		if v, ok := flag[i*10+j]; ok {
			res = append(res, v)
			break
		}
	}

	for i, j := king[0]+1, king[1]-1; i < 8 && j >= 0; i, j = i+1, j-1 {
		if v, ok := flag[i*10+j]; ok {
			res = append(res, v)
			break
		}
	}
	for i, j := king[0]-1, king[1]+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if v, ok := flag[i*10+j]; ok {
			res = append(res, v)
			break
		}
	}
	return res
}
