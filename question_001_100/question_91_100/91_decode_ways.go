package question_91_100

// 91. 解码方法
// https://leetcode-cn.com/problems/decode-ways
// Topics: 字符串 动态规划

func numDecodings(s string) int {
	if len(s) == 0 || (len(s) > 0 && s[0] == '0') {
		return 0
	}
	var flag = make([]int, len(s)+1)
	flag[0], flag[1] = 1, 1
	for i := 1; i < len(s); i++ {
		n := (s[i-1]-'0')*10 + s[i] - '0'
		if n == 0 || (n > 20 && n%10 == 0) {
			return 0
		} else if n == 10 || n == 20 {
			flag[i+1] = flag[i-1]
		} else if (n > 0 && n < 10) || n > 26 {
			flag[i+1] = flag[i]
		} else if n > 10 && n < 27 {
			flag[i+1] = flag[i] + flag[i-1]
		}
	}
	return flag[len(s)]
}
