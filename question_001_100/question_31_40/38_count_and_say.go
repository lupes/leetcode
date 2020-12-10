package question_31_40

import (
	"strconv"
)

// 38. 报数
// https://leetcode-cn.com/problems/count-and-say
// Topics: 字符串

func countAndSay(n int) string {
	result := "1 "
	for i := 2; i <= n; i++ {
		n, c, tmp := 1, result[0], ""
		for j := 1; j < len(result); j++ {
			if c == result[j] {
				n++
			} else {
				tmp += strconv.Itoa(n) + string(c)
				n, c = 1, result[j]
			}
		}
		result = tmp + " "
	}
	return result[:len(result)-1]
}
