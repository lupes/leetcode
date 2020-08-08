package question_841_850

// 848. 字母移位
// https://leetcode-cn.com/problems/shifting-letters
// Topics: 字符串

func shiftingLetters(S string, shifts []int) string {
	res := []byte(S)
	for i := len(shifts) - 1; i >= 0; i-- {
		if i+1 < len(shifts) {
			shifts[i] += shifts[i+1]
		}
		shifts[i] %= 26
		res[i] = (res[i]-'a'+byte(shifts[i]))%26 + 'a'
	}
	return string(res)
}
