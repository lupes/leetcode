package question_01_10

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 3}, []int{}, 2},
		{[]int{1}, []int{}, 1},
		{[]int{}, []int{1}, 1},
		{[]int{}, []int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6}, 3},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
