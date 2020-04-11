package question_971_980

// 977. 有序数组的平方
// https://leetcode-cn.com/problems/squares-of-a-sorted-array
// Topics: 数组 双指针

func sortedSquares(A []int) []int {
	l, r, i := 0, len(A)-1, len(A)-1
	var res = make([]int, len(A))
	for r >= l {
		if A[r]*A[r] > A[l]*A[l] {
			res[i] = A[r] * A[r]
			r--
		} else {
			res[i] = A[l] * A[l]
			l++
		}
		i--
	}
	return res
}
