package question_1171_1180

// 1175. 质数排列
// https://leetcode-cn.com/problems/prime-arrangements
// Topics: 数学

func numPrimeArrangements(n int) int {
	var flag = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	var t, r = 0, 1
	for ; flag[t] <= n; t++ {
	}
	for i := t; i > 0; i-- {
		r = (r * i) % (1e9 + 7)
	}
	for i := n - t; i > 0; i-- {
		r = (r * i) % (1e9 + 7)
	}
	return r
}
