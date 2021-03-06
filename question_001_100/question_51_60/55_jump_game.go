package question_51_60

// 55. 跳跃游戏
// https://leetcode-cn.com/problems/jump-game
// Topics: 贪心算法 数组

func canJump(nums []int) bool {
	max := 0
	for i, n := range nums {
		if max >= i {
			if n+i > max {
				max = n + i
			}
		} else {
			return false
		}
	}
	return true
}
