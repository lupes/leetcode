package question_831_840

// 831. 隐藏个人信息
// https://leetcode-cn.com/problems/masking-personal-information
// Topics: 字符串

func maskPII(S string) string {
	var tmp, isEmail, index = "", false, 0
	for _, c := range S {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.' {
			tmp += string(c)
		} else if c >= 'A' && c <= 'Z' {
			tmp += string(c - 'A' + 'a')
		} else if c == '@' {
			tmp += "@"
			isEmail = true
			index = len(tmp)
		}
	}
	var res = ""
	if isEmail {
		res = tmp[0:1] + "*****" + tmp[index-2:]
	} else {
		switch len(tmp) {
		case 13:
			res = "+***-***-***-"
		case 12:
			res = "+**-***-***-"
		case 11:
			res = "+*-***-***-"
		case 10:
			res = "***-***-"
		}
		res += tmp[len(tmp)-4:]
	}
	return res
}
