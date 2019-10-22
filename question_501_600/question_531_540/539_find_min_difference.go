package question_531_540

import "sort"

// 539. 最小时间差
// https://leetcode-cn.com/problems/mnimum-time-difference/

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	var tmp = make([]int, len(timePoints)+1)
	for i, time := range timePoints {
		t := []byte(time)
		tmp[i] = (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
	}
	tmp[len(timePoints)] = 1440 + tmp[0]
	min := 1440
	for i, _ := range timePoints {
		if min > tmp[i+1]-tmp[i] {
			min = tmp[i+1] - tmp[i]
		}
	}
	return min
}
