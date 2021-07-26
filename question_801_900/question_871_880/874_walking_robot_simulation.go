package question_871_880

import (
	"strconv"
)

// 874. 模拟行走机器人
// https://leetcode-cn.com/problems/walking-robot-simulation
// Topics: 贪心算法

func robotSim(commands []int, obstacles [][]int) int {
	var flag = make(map[string]struct{})
	for _, obstacle := range obstacles {
		flag[strconv.Itoa(obstacle[0])+","+strconv.Itoa(obstacle[1])] = struct{}{}
	}
	var x, y, t, max = 0, 0, 0, 0
	xd, yd := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	for _, command := range commands {
		if command == -1 {
			t = (t + 1) % 4
		} else if command == -2 {
			t = (t + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				if _, ok := flag[strconv.Itoa(x+xd[t])+","+strconv.Itoa(y+yd[t])]; ok {
					break
				} else {
					x, y = x+xd[t], y+yd[t]
				}
			}
		}
		if x*x+y*y > max {
			max = x*x + y*y
		}
	}

	return max
}
