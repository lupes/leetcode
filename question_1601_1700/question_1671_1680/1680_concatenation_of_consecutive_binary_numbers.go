package question_1671_1680

// 1680. 连接连续二进制数字
// https://leetcode-cn.com/problems/concatenation-of-consecutive-binary-numbers/
// Topics: 数学 位运算

func concatenatedBinary(n int) int {
	var shift, res, i, t uint = 0, 0, 1, uint(n)
	for ; i <= t; i++ {
		if i&(i-1) == 0 {
			shift++
		}
		res = (res<<shift + i) % (1e9 + 7)
	}
	return int(res)
}
