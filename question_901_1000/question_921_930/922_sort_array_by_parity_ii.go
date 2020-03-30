package question_921_930

// 922. 按奇偶排序数组 II
// https://leetcode-cn.com/problems/sort-array-by-parity-ii
// Topics: 排序 数组

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A)-1 && j < len(A); {
		if A[i]%2 == 1 && A[j]%2 == 0 {
			A[i], A[j] = A[j], A[i]
		} else if A[i]%2 == 0 {
			i += 2
		} else if A[j]%2 == 1 {
			j += 2
		}
	}
	return A
}
