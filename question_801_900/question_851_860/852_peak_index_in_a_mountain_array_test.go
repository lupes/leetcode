package question_851_860

import (
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 1, 2, 1, 0}, 2},
		{[]int{0, 1, 2, 3, 2, 0}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.A); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
