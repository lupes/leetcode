package question_1311_1320

// 1318. 或运算的最小翻转次数
// https://leetcode-cn.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
// Topics: 位运算

func minFlips(a int, b int, c int) int {
	var res, t, r int
	for a > 0 || b > 0 || c > 0 {
		t, r, a, b, c = a&1+b&1, c&1, a>>1, b>>1, c>>1
		if r == 0 {
			res += t
		} else if r == 1 && t == 0 {
			res += 1
		}
	}
	return res
}
