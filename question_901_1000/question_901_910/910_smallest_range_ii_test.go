package question_901_910

import (
	"testing"
)

func Test_smallestRangeII(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want int
	}{
		{[]int{1}, 0, 0},
		{[]int{0, 10}, 2, 6},
		{[]int{1, 3, 6}, 3, 3},
		{[]int{1, 3, 6, 7}, 3, 3},
		{[]int{7, 8, 8}, 5, 1},
		{[]int{9, 10, 5, 9}, 5, 5},
		{[]int{3, 4, 7, 0}, 5, 7},
		{[]int{4, 8, 2, 7, 2}, 5, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := smallestRangeII(tt.A, tt.K); got != tt.want {
				t.Errorf("smallestRangeII() = %v, want %v", got, tt.want)
			}
		})
	}
}
