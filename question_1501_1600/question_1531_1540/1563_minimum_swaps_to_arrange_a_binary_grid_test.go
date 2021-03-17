package question_1531_1540

import (
	"testing"
)

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}, 3},
		{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}, -1},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}, 0},
		{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSwaps(tt.grid); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
