package question_81_90

// 88. 合并两个有序数组
// https://leetcode-cn.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if n == 0 || (m > 0 && nums1[m-1] >= nums2[n-1]) {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
