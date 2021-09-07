package question_1221_1230

// 1221. 分割平衡字符串
// https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
// Topics: 贪心算法 字符串

func balancedStringSplit(s string) int {
	var t, r = 0, 0
	for _, c := range s {
		switch c {
		case 'L':
			t++
		case 'R':
			t--
		}
		if t == 0 {
			r++
		}
	}
	return r
}
