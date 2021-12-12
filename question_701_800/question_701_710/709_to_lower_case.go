package question_701_710

// 709. 转换成小写字母
// https://leetcode-cn.com/problems/to-lower-case
// Topics: 字符串

func toLowerCase(str string) string {
	var bytes = []byte(str)
	for i, b := range bytes {
		if b >= 'A' && b <= 'Z' {
			bytes[i] = b + 'a' - 'A'
		}
	}
	return string(bytes)
}
