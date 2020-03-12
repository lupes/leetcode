package question_901_910

// 905. 按奇偶排序数组
// https://leetcode-cn.com/problems/sort-array-by-parity
// Topics: 数组

func sortArrayByParity(A []int) []int {
	for left, right := 0, len(A)-1; right > left; {
		if A[left]%2 == 0 {
			left++
		} else if A[right]%2 == 1 {
			right--
		} else {
			A[left], A[right] = A[right], A[left]
		}
	}
	return A
}
