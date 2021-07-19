package question_821_830

// 822. 翻转卡片游戏
// https://leetcode-cn.com/problems/card-flipping-game
// Topics:

func flipgame(fronts []int, backs []int) int {
	var flag = make([]int, 2000+1)
	for i := range fronts {
		if fronts[i] == backs[i] {
			flag[fronts[i]] = -1
		} else {
			if flag[fronts[i]] != -1 {
				flag[fronts[i]]++
			}
			if flag[backs[i]] != -1 {
				flag[backs[i]]++
			}
		}
	}
	for i, n := range flag {
		if n > 0 {
			return i
		}
	}
	return 0
}
