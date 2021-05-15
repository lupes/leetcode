package question_1681_1690

// 1688. 比赛中的配对次数
// https://leetcode-cn.com/problems/count-of-matches-in-tournament/
// Topics: 回溯算法

func numberOfMatches(n int) int {
	var res = 0
	for n > 1 {
		if n&1 == 1 {
			res, n = res+(n-1)>>1, (n-1)>>1+1
		} else {
			res, n = res+n>>1, n>>1
		}
	}
	return res
}
