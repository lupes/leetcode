package question_941_950

import (
	"sort"
)

// 949. 给定数字能组成的最大时间
// https://leetcode-cn.com/problems/largest-time-for-given-digits
// Topics: 数学

func largestTimeFromDigits(A []int) string {
	var res = largestTimeFromDigitsHelper(nil, A, 1)
	if len(res) == 0 {
		return ""
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return res[0]
}

func largestTimeFromDigitsHelper(flag []byte, A []int, i int) []string {
	var res []string
	for j, n := range A {
		if n == -1 {
			continue
		}
		if i == 1 && n > 2 {
			continue
		}
		if i == 2 && flag[0] == 2 && n > 3 {
			continue
		}
		if i == 3 && n > 5 {
			continue
		}
		if i == 4 {
			res = append(res, string([]byte{flag[0] + '0', flag[1] + '0', ':', flag[2] + '0', byte(n) + '0'}))
		} else {
			t := A[j]
			A[j] = -1
			tmp := largestTimeFromDigitsHelper(append(flag, byte(n)), A, i+1)
			res = append(res, tmp...)
			A[j] = t
		}
	}
	return res
}
