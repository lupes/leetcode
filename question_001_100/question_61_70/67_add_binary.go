package question_61_70

import (
	"strconv"
)

// 67. 二进制求和
// https://leetcode-cn.com/problems/add-binary

func addBinary(a string, b string) string {
	ia := len(a) - 1
	ib := len(b) - 1
	res := ""
	var h uint8
	for ia >= 0 || ib >= 0 {
		var ta, tb uint8
		if ia >= 0 {
			ta = a[ia] - '0'
		}
		if ib >= 0 {
			tb = b[ib] - '0'
		}
		tc := ta + tb + h
		h = tc / 2
		res = strconv.Itoa(int(tc%2)) + res
		ia--
		ib--
	}
	if h == 1 {
		res = "1" + res
	}
	return res
}
