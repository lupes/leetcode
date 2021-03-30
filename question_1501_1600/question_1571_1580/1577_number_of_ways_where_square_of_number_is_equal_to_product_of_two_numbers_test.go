package question_1571_1580

import (
	"testing"
)

func Test_numTriplets(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{7, 4}, []int{5, 2, 8, 9}, 1},
		{[]int{1, 1}, []int{1, 1, 1}, 9},
		{[]int{7, 7, 8, 3}, []int{1, 2, 9, 7}, 2},
		{[]int{4, 7, 9, 11, 23}, []int{3, 5, 1024, 12, 18}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numTriplets(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("numTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
