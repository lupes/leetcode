package question_851_860

// 852. 山脉数组的峰顶索引
// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array
// Topics: 二分查找

func peakIndexInMountainArray(A []int) int {
	left, right, center := 0, len(A), 0
	for right > left {
		center = (left + right) / 2
		if A[center] > A[center-1] && A[center] > A[center+1] {
			return center
		} else if A[center] > A[center-1] && A[center] < A[center+1] {
			left = center + 1
		} else {
			right = center
		}
	}
	return 0
}
