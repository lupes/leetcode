package question_941_950

// 941. 有效的山脉数组
// https://leetcode-cn.com/problems/valid-mountain-array
// Topics: 数组

func validMountainArray(A []int) bool {
	var i = 0
	for i = 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return false
		}
		if A[i] > A[i+1] {
			break
		}
	}
	if i == len(A)-1 || i == 0 {
		return false
	}
	for ; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			return false
		}
	}
	return true
}
