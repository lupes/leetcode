package question_1001_1010

// 1009. 十进制整数的反码
// https://leetcode-cn.com/problems/complement-of-base-10-integer
// Topics: 数学

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	for i := 0; ; i++ {
		if 1<<i > N {
			return 1<<i - 1 - N
		}
	}
}
