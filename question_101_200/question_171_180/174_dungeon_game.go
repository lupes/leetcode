package question_171_180

// 174. 地下城游戏
// https://leetcode-cn.com/problems/dungeon-game
// Topics: 二分查找 动态规划

func calculateMinimumHP(dungeon [][]int) int {
	rl, cl := len(dungeon), len(dungeon[0])
	var res = make([][]int, rl)
	for i := range res {
		res[i] = make([]int, cl)
	}
	for i := rl - 1; i >= 0; i-- {
		for j := cl - 1; j >= 0; j-- {
			if i == rl-1 && j == cl-1 {
				if dungeon[i][j] < 1 {
					res[i][j] = 1 - dungeon[i][j]
				} else {
					res[i][j] = 1
				}
			} else if i == rl-1 {
				tmp := dungeon[i][j] - res[i][j+1]
				if tmp >= 0 {
					res[i][j] = 1
				} else {
					res[i][j] = -tmp
				}
			} else if j == cl-1 {
				tmp := dungeon[i][j] - res[i+1][j]
				if tmp >= 0 {
					res[i][j] = 1
				} else {
					res[i][j] = -tmp
				}
			} else if i < rl-1 && j < cl-1 {
				tmp := dungeon[i][j] - res[i][j+1]
				if tmp >= 0 {
					res[i][j] = 1
				} else {
					res[i][j] = -tmp
				}
				tmp = dungeon[i][j] - res[i+1][j]
				if tmp >= 0 {
					tmp = 1
				} else {
					tmp = -tmp
				}
				if res[i][j] > tmp {
					res[i][j] = tmp
				}
			}
		}
	}
	return res[0][0]
}
