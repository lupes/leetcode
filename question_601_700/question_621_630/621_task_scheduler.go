package question_621_630

import (
	"sort"
)

// 621. 任务调度器
// https://leetcode-cn.com/problems/task-scheduler
// Topics: 贪心算法 队列 数组

func leastInterval(tasks []byte, n int) int {
	var ts = make([]int, 26)
	for _, t := range tasks {
		ts[t-'A']++
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i] > ts[j]
	})
	var max = 0
	for i := 0; i < 26 && ts[i] == ts[0]; i++ {
		max++
	}
	res := (ts[0]-1)*(n+1) + max
	if len(tasks) > res {
		return len(tasks)
	}
	return res
}
