package question_91_100

// 91. 解码方法
// https://leetcode-cn.com/problems/decode-ways

func numDecodings(s string) int {
	if (len(s) == 1 && s[0] == '0') || (len(s) == 2 && s[0] > '2' && s[1] == '0') {
		return 0
	}
	if len(s) < 2 {
		return 1
	}

	if (s[0] > '2' && s[1] == '0') || s[0] == '0' {
		return 0
	} else if s[0] > '2' || (len(s) > 2 && s[2] == '0') {
		return numDecodings(s[1:])
	} else if (s[0] == '2' && s[1] > '6') || s[1] == '0' {
		return numDecodings(s[2:])
	} else {
		n1 := numDecodings(s[1:])
		n2 := numDecodings(s[2:])
		if n1 == 0 || n2 == 0 {
			return 0
		}
		return n1 + n2
	}
}
