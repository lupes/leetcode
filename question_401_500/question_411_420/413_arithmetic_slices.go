package question_411_420

// 413. 等差数列划分
// https://leetcode-cn.com/problems/arithmetic-slices
// Topics: 数学 动态规划

func numberOfArithmeticSlices(A []int) int {
	var tmp, res = make([]int, len(A)), 0
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			tmp[i] = tmp[i-1] + 1
		}
		res += tmp[i]
	}
	return res
}
