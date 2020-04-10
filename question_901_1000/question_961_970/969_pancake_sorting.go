package question_961_970

// 969. 煎饼排序
// https://leetcode-cn.com/problems/pancake-sorting
// Topics: 排序 数组

func pancakeSort(A []int) []int {
	var res []int
	var l = len(A)
	for i := 0; i < l; i++ {
		var max, index = A[0], 0
		for j := 1; j < l-i; j++ {
			if A[j] > max {
				max, index = A[j], j
			}
		}
		if index != l-i-1 {
			pancakeSortHelper(A, index)
			pancakeSortHelper(A, l-i-1)
			res = append(res, index+1, l-i)
		}
	}
	return res
}

func pancakeSortHelper(A []int, k int) {
	for i := 0; i <= k/2; i++ {
		A[i], A[k-i] = A[k-i], A[i]
	}
}
