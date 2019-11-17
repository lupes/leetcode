package question_391_400

import (
	"strconv"
	"strings"
)

// 394. 字符串解码
// https://leetcode-cn.com/problems/decode-string/
// Topics: 栈 深度优先搜索

func decodeString(s string) string {
	var res, n, j, k, l = "", 0, 0, 0, len(s)
	for i := 0; i < l; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			res += string(s[i])
		} else if s[i] >= '1' && s[i] <= '9' {
			for j = i + 1; j < len(s); j++ {
				if s[j] == '[' {
					n, _ = strconv.Atoi(s[i:j])
					break
				}
			}
			var t = 1
			for k = j + 1; k < len(s) && t > 0; k++ {
				if s[k] == '[' {
					t++
				} else if s[k] == ']' {
					t--
				}
			}
			res += strings.Repeat(decodeString(s[j+1:k-1]), n)
			i = k - 1
		}
	}
	return res
}
