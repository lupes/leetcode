package question_1361_1370

// 1376. 通知所有员工所需的时间
// https://leetcode-cn.com/problems/time-needed-to-inform-all-employees/
// Topics: 深度优先搜索

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var res = 0
	var times = make([]int, len(manager))
	for i := range times {
		times[i] = -1
	}
	for i, _ := range manager {
		if times[i] != -1 {
			continue
		}
		t := numOfMinutesHelper(i, manager, informTime, times)
		if t > res {
			res = t
		}
	}
	return res
}

func numOfMinutesHelper(nowId int, manager []int, informTime []int, times []int) int {
	if manager[nowId] == -1 {
		return informTime[nowId]
	} else if times[nowId] == -1 {
		times[nowId] = numOfMinutesHelper(manager[nowId], manager, informTime, times)
	}

	return informTime[nowId] + times[nowId]
}
