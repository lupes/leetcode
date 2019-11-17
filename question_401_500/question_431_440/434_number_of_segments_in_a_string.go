package question_431_440

// 434. 字符串中的单词数
// https://leetcode-cn.com/problems/number-of-segments-in-a-string/

func countSegments(s string) int {
	var res, flag = 0, false
	for _, n := range s {
		if n == ' ' && flag {
			res++
			flag = false
		} else if n != ' ' {
			flag = true
		}
	}
	if flag {
		res++
	}
	return res
}
