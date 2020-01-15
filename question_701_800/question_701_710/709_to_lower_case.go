package question_701_710

// 709. 转换成小写字母
// https://leetcode-cn.com/problems/to-lower-case
// Topics: 字符串

func toLowerCase(str string) string {
	var res = ""
	for _, n := range str {
		if n >= 'A' && n <= 'Z' {
			n += 'a' - 'A'
		}
		res += string(n)
	}
	return res
}
