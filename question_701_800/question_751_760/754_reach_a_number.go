package question_751_760

// 754. 到达终点数字
// https://leetcode-cn.com/problems/reach-a-number
// Topics: 数学

func reachNumber(target int) int {
	if target < 0 {
		target *= -1
	}
	var tmp = 0
	for i := 1; ; i++ {
		tmp += i
		if tmp == target || (tmp > target && (tmp-target)%2 == 0) {
			return i
		}
	}
}
