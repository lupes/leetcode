package question_881_890

import (
	"testing"
)

func Test_projectionArea(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2}}, 5},
		{[][]int{{1, 0}, {0, 2}}, 8},
		{[][]int{{1, 2}, {3, 4}}, 17},
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 14},
		{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}, 21},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := projectionArea(tt.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
