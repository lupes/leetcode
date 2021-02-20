package question_1451_1460

import (
	"testing"
)

func Test_maxDotProduct(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{2, 1, -2, 5}, []int{3, 0, -6}, 18},
		{[]int{3, -2}, []int{2, -6, 7}, 21},
		{[]int{-1, -1}, []int{1, 1}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxDotProduct(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
