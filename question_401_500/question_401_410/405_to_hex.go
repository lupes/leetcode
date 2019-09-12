package question_401_410

// 405. 数字转换为十六进制数
// https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal/

var flags = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

func toHex(num int) string {
	if num < 0 {
		num = 0xffffffff + num + 1
	} else if num == 0 {
		return "0"
	}
	var res = ""
	for num > 0 {
		res = flags[num%16] + res
		num = num / 16
	}
	return res
}
