package question_441_450

// 443. 压缩字符串
// https://leetcode-cn.com/problems/string-compression/
// Topics: 字符串

func compress(chars []byte) int {
	n, j := 0, 0
	for i := 0; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[j] {
			n++
		} else {
			a, b, c, d := n%10000/1000+'0', n%1000/100+'0', n%100/10+'0', n%10+'0'
			if n > 999 {
				j++
				chars[j] = byte(a)
			}
			if n > 99 {
				j++
				chars[j] = byte(b)
			}
			if n > 9 {
				j++
				chars[j] = byte(c)
			}
			if n > 1 {
				j++
				chars[j] = byte(d)
			}
			if i < len(chars) {
				j++
				chars[j] = chars[i]
			}
			n = 1
		}
	}
	return j + 1
}
