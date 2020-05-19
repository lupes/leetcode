package question_1011_1020

// 1017. 负二进制转换
// https://leetcode-cn.com/problems/convert-to-base-2
// Topics: 数学

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}
	var res, r = "", 0
	for N != 0 {
		r, N = N%-2, N/-2
		if r == -1 {
			N, r = N+1, 1
		}
		res = string('0'+r) + res
	}
	return res
}
