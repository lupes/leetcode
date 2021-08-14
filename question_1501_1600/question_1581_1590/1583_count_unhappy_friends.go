package question_1581_1590

// 1583. 统计不开心的朋友
// https://leetcode-cn.com/problems/count-unhappy-friends/
// Topics: 数组 模拟

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	var res, flag, friend = 0, make([][]int, n), make([]int, n)
	for i, preference := range preferences {
		flag[i] = make([]int, n)
		flag[i][i] = n + 1
		for j, n := range preference {
			flag[i][n] = j
		}
	}

	for _, pair := range pairs {
		friend[pair[0]] = pair[1]
		friend[pair[1]] = pair[0]
	}

	for x, y := range friend {
		for i, f := range flag[x] {
			if f < flag[x][y] && flag[i][x] < flag[i][friend[i]] {
				res++
				break
			}
		}
	}

	return res
}
