package question_441_450

// 441. 排列硬币
// https://leetcode-cn.com/problems/arranging-coins/

func arrangeCoins(n int) int {
	l, r, c := 0, n, 0
	for r > l {
		c = (l+r)/2 + 1
		t := (1 + c) * c / 2
		if t == n {
			return c
		} else if t < n {
			l = c
		} else if t > n {
			r = c - 1
		}
	}
	return l
}
