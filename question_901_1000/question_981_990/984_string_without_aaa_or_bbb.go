package question_981_990

// 984. 不含 AAA 或 BBB 的字符串
// https://leetcode-cn.com/problems/string-without-aaa-or-bbb
// Topics: 贪心算法

func strWithout3a3b(A int, B int) string {
	x, y, xnum, ynum, res := "a", "b", A, B, ""
	if B > A {
		x, y, xnum, ynum = "b", "a", B, A
	}
	for xnum > 0 && ynum >= 0 {
		if xnum > 0 && ynum == 0 {
			res, xnum = res+x, xnum-1
		} else if xnum > ynum {
			res, xnum, ynum = res+x+x+y, xnum-2, ynum-1
		} else {
			res, xnum, ynum = res+x+y, xnum-1, ynum-1
		}
	}
	return res
}
