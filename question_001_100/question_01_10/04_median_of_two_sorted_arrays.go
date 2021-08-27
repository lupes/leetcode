package question_01_10

// 4. 寻找两个有序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// Topics: 数组 二分查找 分治算法

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2, l := len(nums1), len(nums2), len(nums1)+len(nums2)

	a, b, t := 0, 0, 0
	for i1, i2, i3 := 0, 0, 0; i3 < l; i3++ {
		if l1 == i1 || (l1 > i1 && l2 > i2 && nums1[i1] >= nums2[i2]) {
			t = nums2[i2]
			i2++
		} else if l2 == i2 || (l1 > i1 && l2 > i2 && nums1[i1] < nums2[i2]) {
			t = nums1[i1]
			i1++
		}
		if i3 == l/2-1 {
			a = t
		} else if i3 == l/2 {
			b = t
			break
		}
	}
	if l&1 == 0 {
		return float64(a+b) / 2
	} else {
		return float64(b)
	}
}
