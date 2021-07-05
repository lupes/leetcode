package question_661_670

// 667. 优美的排列 II
// https://leetcode-cn.com/problems/beautiful-arrangement-ii
// Topics: 数组

func constructArray(n int, k int) []int {
	var res = make([]int, 0, n)
	for a, b := 1, 1+k; b > a; a, b = a+1, b-1 {
		res = append(res, a, b)
	}
	if k%2 == 0 {
		res = append(res, k/2+1)
	}
	for j := k + 2; j <= n; j++ {
		res = append(res, j)
	}
	return res
}
