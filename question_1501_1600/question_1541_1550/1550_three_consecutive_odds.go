package question_1541_1550

// 1550. 存在连续三个奇数的数组
// https://leetcode-cn.com/problems/three-consecutive-odds/
// Topics: 数组

func threeConsecutiveOdds(arr []int) bool {
	var dp = 0
	for _, n := range arr {
		if n%2 == 0 {
			dp = 0
		} else {
			dp += 1
		}
		if dp == 3 {
			return true
		}
	}
	return false
}
