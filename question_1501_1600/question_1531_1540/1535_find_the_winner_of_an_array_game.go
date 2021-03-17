package question_1531_1540

// 1535. 找出数组游戏的赢家
// https://leetcode-cn.com/problems/find-the-winner-of-an-array-game/
// Topics: 数组

func getWinner(arr []int, k int) int {
	var res, num = arr[0], 0
	for i := 1; i < len(arr) && num < k; i++ {
		if res > arr[i] {
			num += 1
		} else {
			res, num = arr[i], 1
		}
	}
	return res
}
