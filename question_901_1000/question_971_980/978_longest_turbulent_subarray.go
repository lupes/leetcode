package question_971_980

// 978. 最长湍流子数组
// https://leetcode-cn.com/problems/longest-turbulent-subarray
// Topics: 数组 动态规划 Sliding Window

func maxTurbulenceSize(A []int) int {
	var flag, res, tmp = 0, 1, 1
	for i := 1; i < len(A); i++ {
		if (A[i-1] > A[i] && flag == -1) || (A[i-1] < A[i] && flag == 1) {
			tmp++
		} else if A[i-1] == A[i] {
			tmp = 1
		} else {
			tmp = 2
		}
		if tmp > res {
			res = tmp
		}
		if A[i-1] > A[i] {
			flag = 1
		} else if A[i-1] < A[i] {
			flag = -1
		} else {
			flag = 0
		}
	}
	return res
}
