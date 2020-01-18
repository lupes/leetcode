package question_711_720

import (
	"testing"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{10, 5, 2, 6}, 10, 3},
		{[]int{1, 2, 3}, 0, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.nums, tt.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
