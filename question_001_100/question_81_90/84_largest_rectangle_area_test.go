package question_81_90

import (
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 2}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{2, 1, 2}, 3},
		{[]int{1, 1, 2}, 3},
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 1, 5, 6, 2, 1000000}, 1000000},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
