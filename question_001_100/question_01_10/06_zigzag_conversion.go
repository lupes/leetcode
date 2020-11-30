package question_01_10

// 6. Z 字形变换
// https://leetcode-cn.com/problems/zigzag-conversion
// Topics: 字符串

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) == 0 {
		return s
	}
	res, num := make([]byte, 0, len(s)), numRows-1
	for i := 0; i < len(s); i += num * 2 {
		res = append(res, s[i])
	}
	for i := 1; i < num; i++ {
		for j := 0; j < len(s)/num+1; j++ {
			if j%2 == 0 && num*j+i < len(s) {
				res = append(res, s[num*j+i])
			} else if j%2 == 1 && num*j+(num-i) < len(s) {
				res = append(res, s[num*j+(num-i)])
			}
		}
	}
	for i := num; i < len(s); i += num * 2 {
		res = append(res, s[i])
	}

	return string(res)
}
