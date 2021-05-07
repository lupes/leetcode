package question_1651_1660

// 1654. 到家的最少跳跃次数
// https://leetcode-cn.com/problems/minimum-jumps-to-reach-home/
// Topics: 广度优先搜索 动态规划

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}

	var flag = make([]int, x+b*6)
	for _, f := range forbidden {
		if f < x+b*6 {
			flag[f] = 1
		}
	}

	var now = map[int]bool{0: true}
	for num := 1; len(now) > 0; num++ {
		var next = make(map[int]bool)
		for t, f := range now {
			front, back := t+a, t-b
			if !f {
				front, back = t+a, t
			}

			if front == x || back == x {
				return num
			}

			if front < len(flag) && flag[front] <= 0 {
				next[front] = true
				flag[front] = num
			}

			if back > 0 && flag[back] == 0 {
				next[back] = false
				flag[back] = -num
			}
		}
		now = next
	}
	return -1
}
