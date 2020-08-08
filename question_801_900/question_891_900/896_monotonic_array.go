package question_891_900

// 896. 单调数列
// https://leetcode-cn.com/problems/monotonic-array
// Topics: 数组

func isMonotonic(A []int) bool {
	var flag = 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] || (A[i] > A[i-1] && flag == 1) || (A[i] < A[i-1] && flag == -1) {
			continue
		} else if (A[i] > A[i-1] && flag == -1) || (A[i] < A[i-1] && flag == 1) {
			return false
		} else if A[i] > A[i-1] && flag == 0 {
			flag = 1
		} else if A[i] < A[i-1] && flag == 0 {
			flag = -1
		}
	}
	return true
}
