package question_491_500

// 495. 提莫攻击
// https://leetcode-cn.com/problems/teemo-attacking/

func findPoisonedDuration(timeSeries []int, duration int) int {
	res, last := 0, -duration
	for _, t := range timeSeries {
		res += duration
		if last+duration > t {
			res -= duration - (t - last)
		}
		last = t
	}
	return res
}
