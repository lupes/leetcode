package question_1101_1110

import (
	"strings"
)

// 1108. IP 地址无效化
// https://leetcode-cn.com/problems/defanging-an-ip-address
// Topics: 字符串

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", 3)
}
