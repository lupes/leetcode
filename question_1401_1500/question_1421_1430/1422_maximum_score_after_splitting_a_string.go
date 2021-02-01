package question_1421_1430

// 1422. 分割字符串的最大得分
// https://leetcode-cn.com/problems/maximum-score-after-splitting-a-string/
// Topics: 字符串

func maxScore1422(s string) int {
	max, leftZero, leftOne, rightZero, rightOne := 0, 0, 0, 0, 0
	for _, n := range s {
		if n == '0' {
			rightZero++
		} else {
			rightOne++
		}
	}

	for _, n := range s[:len(s)-1] {
		if n == '0' {
			leftZero++
			rightZero--
		} else {
			leftOne++
			rightOne--
		}
		if leftZero+rightOne > max {
			max = leftZero + rightOne
		}
	}
	return max
}
