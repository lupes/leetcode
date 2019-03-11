package question_01_10

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)
	size := len(nums3)
	rem, quo := size%2, size/2
	median := 0.0
	if rem == 0 {
		median = float64(nums3[quo-1]+nums3[quo]) / 2
	} else {
		median = float64(nums3[quo])
	}
	return median
}
