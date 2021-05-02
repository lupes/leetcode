package question_1661_1670

// 1665. 完成所有任务的最少初始能量
// https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/
// Topics: 贪心算法

import (
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	var res, now = 0, 0
	for _, task := range tasks {
		if now < task[1] {
			res, now = res+task[1]-now, task[1]
		}
		now -= task[0]
	}
	return res
}
