package question_921_930

// 926. 将字符串翻转到单调递增
// https://leetcode-cn.com/problems/flip-string-to-monotone-increasing
// Topics: 数组

func minFlipsMonoIncr(S string) int {
	var min, zero, one = len(S), 0, 0
	for _, c := range S {
		if c == '0' {
			zero++
		}
	}
	if zero == 0 || zero == len(S) {
		return 0
	}
	min = zero
	for _, c := range S {
		if c == '1' {
			one++
		} else {
			zero--
		}
		if one+zero < min {
			min = one + zero
		}
		if min == 0 {
			return 0
		}
	}
	return min
}
