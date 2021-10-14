package offer

import (
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 1}, 1},
		{[]int{3, 2, 1}, 0},
		{[]int{1, 3, 5, 4, 2}, 2},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
		{[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
