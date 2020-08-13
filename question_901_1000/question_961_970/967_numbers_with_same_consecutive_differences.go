package question_961_970

// 967. 连续差相同的数字
// https://leetcode-cn.com/problems/numbers-with-same-consecutive-differences
// Topics: 动态规划

func numsSameConsecDiff(N int, K int) []int {
	var res = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := N; i > 1; i-- {
		var next []int
		for _, n := range res {
			if n == 0 {
				continue
			}
			t := n % 10
			if t-K >= 0 {
				next = append(next, n*10+t-K)
			}
			if K != 0 && t+K < 10 {
				next = append(next, n*10+t+K)
			}
		}
		res = next
	}
	return res
}
