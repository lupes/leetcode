package question_331_340

// 338. 比特位计数
// https://leetcode-cn.com/problems/counting-bits/
// Topics: 位运算 动态规划

func countBits(num int) []int {
	var res []int
	for i := 0; i <= num; i++ {
		t, n := 0, i
		for n > 0 {
			t += int(n & 1)
			n >>= 1
		}
		res = append(res, t)
	}
	return res
}
