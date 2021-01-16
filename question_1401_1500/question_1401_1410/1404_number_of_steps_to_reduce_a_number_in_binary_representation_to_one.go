package question_1401_1410

// 1404. 将二进制表示减到 1 的步骤数
// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
// Topics: 位运算 字符串

func numSteps(s string) int {
	var res, tmp, c = 0, uint8(0), uint8(0)
	for len(s) > 1 {
		c, s = s[len(s)-1]-'0', s[:len(s)-1]
		if c+tmp == 1 {
			tmp = 1
			res++
		}
		res++
	}

	return res + int((s[0]-'0')&tmp)
}
