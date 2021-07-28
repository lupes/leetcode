package question_891_900

import (
	"bytes"
)

// 899. 有序队列
// https://leetcode-cn.com/problems/orderly-queue
// Topics: 数学 字符串

func orderlyQueue(S string, K int) string {
	if K == 1 {
		var buffer, res = bytes.Buffer{}, S
		for 0 < len(S) {
			buffer.WriteByte(S[0])
			S = S[1:]
			if res > S+buffer.String() {
				res = S + buffer.String()
			}
		}
		return res
	}

	var flag = [26]int{}
	for _, c := range S {
		flag[c-'a'] += 1
	}

	var buffer = bytes.Buffer{}
	for i, n := range flag {
		for j := 0; j < n; j++ {
			buffer.WriteByte(byte(i) + 'a')
		}
	}

	return buffer.String()
}
