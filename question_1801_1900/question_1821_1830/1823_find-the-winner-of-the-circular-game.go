package question_1821_1830

// 1823. 找出游戏的获胜者
// https://leetcode.cn/problems/find-the-winner-of-the-circular-game/
// 递归 队列 数组 数学 模拟

func findTheWinner(n int, k int) int {
	var queue = make([]int, 0, n)
	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}

	for start, i := 0, 0; len(queue) > 1; start = i {
		i = (k+start)%len(queue) - 1
		if i == 0 {
			queue = queue[1:]
		} else if i == -1 {
			queue = queue[:len(queue)-1]
			i = 0
		} else {
			queue = append(queue[:i], queue[i+1:]...)
		}
	}
	return queue[0]
}
