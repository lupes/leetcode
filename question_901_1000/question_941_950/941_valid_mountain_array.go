package question_941_950

// 941. 有效的山脉数组
// https://leetcode-cn.com/problems/valid-mountain-array
// Topics: 数组

func validMountainArray(A []int) bool {
	l := len(A)
	for i := 0; i < l; i++ {

	}
	for left, right := 0, l; right >= left; {
		if right > 0 && A[right] < A[right-1] {
			right--
		}
	}
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
