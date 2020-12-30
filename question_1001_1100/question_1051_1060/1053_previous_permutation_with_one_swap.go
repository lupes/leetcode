package question_1051_1060

// 1053. 交换一次的先前排列
// https://leetcode-cn.com/problems/previous-permutation-with-one-swap
// Topics: 贪心算法 数组

func prevPermOpt1(A []int) []int {
	var t, k int
	for i := len(A) - 2; i >= 0; i-- {
		t, k = 0, 0
		for j := i + 1; j < len(A) && A[j] < A[i]; j++ {
			if A[i] > A[j] && A[j] > t {
				t, k = A[j], j
			}
		}
		if t > 0 {
			A[i], A[k] = A[k], A[i]
			return A
		}
	}
	return A
}
