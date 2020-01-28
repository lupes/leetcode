package question_731_740

// 739. 每日温度
// https://leetcode-cn.com/problems/daily-temperatures
// Topics: 栈 哈希表

func dailyTemperatures(T []int) []int {
	var res = make([]int, len(T))
	for i, t := range T {
		for j := i + 1; j < len(T); j++ {
			if T[j] > t {
				res[i] = j - i
				break
			}
		}
	}
	return res
}
