package question_391_400

// 393. UTF-8 编码验证
// https://leetcode-cn.com/problems/utf-8-validation/
// Topics: 位运算

func validUtf8(data []int) bool {
	l, n := len(data), 0
	for i := 0; i < l; i++ {
		if data[i]>>7 != 0 {
			switch {
			case data[i]>>3 == 30:
				n = 3
			case data[i]>>4 == 14:
				n = 2
			case data[i]>>5 == 6:
				n = 1
			default:
				return false
			}
			if i+n >= l {
				return false
			}
			for k := 1; k <= n; k++ {
				if data[k+i]>>6 != 2 {
					return false
				}
			}
			i += n
		}
	}
	return true
}
