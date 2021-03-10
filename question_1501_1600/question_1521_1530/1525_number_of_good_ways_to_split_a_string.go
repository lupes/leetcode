package question_1521_1530

// 1525. 字符串的好分割数目
// https://leetcode-cn.com/problems/number-of-good-ways-to-split-a-string/
// Topics: 字符串 位运算

func numSplits(s string) int {
	var left, right = make(map[int32]int, 26), make(map[int32]int, 26)
	for _, c := range s {
		right[c]++
	}
	var res = 0
	for _, c := range s {
		left[c]++
		right[c]--
		if right[c] == 0 {
			delete(right, c)
		}
		if len(left) == len(right) {
			res++
		}
	}
	return res
}
