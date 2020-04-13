package question_981_990

// 985. 查询后的偶数和
// https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries
// Topics: 数组

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var sum = 0
	for _, n := range A {
		if n%2 == 0 {
			sum += n
		}
	}
	var res []int
	for _, query := range queries {
		if A[query[1]]%2 == 0 {
			sum -= A[query[1]]
		}
		if (A[query[1]]+query[0])%2 == 0 {
			sum += A[query[1]] + query[0]
		}
		A[query[1]] = A[query[1]] + query[0]
		res = append(res, sum)
	}
	return res
}
