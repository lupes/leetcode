package question_381_390

// 390. 消除游戏
// https://leetcode-cn.com/problems/elimination-game/
// Topics:

func lastRemaining(n int) int {
	switch n {
	case 1:
		return 1
	default:
		return 2 * (n/2 + 1 - lastRemaining(n/2))
	}
}
