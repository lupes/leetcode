package question_531_540

// 539. 最小时间差
// https://leetcode-cn.com/problems/mnimum-time-difference/
// Topics: 字符串

func findMinDifference(timePoints []string) int {
	var times = make([]bool, 1440)
	var left, min = 1441, 1441
	for _, time := range timePoints {
		i := int((time[0]-'0')*10+time[1]-'0')*60 + int(time[3]-'0')*10 + int(time[4]-'0')
		if times[i] {
			return 0
		}
		times[i] = true
		if i < left {
			left = i
		}
	}
	times = append(times[left:], times[:left+1]...)

	left = 0
	for i := left + 1; i < len(times); i++ {
		if times[i] {
			if i-left < min {
				min = i - left
			}
			left = i
		}
	}
	return min
}
