package question_1091_1100

import (
	"testing"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 1}, {1, 0}}, 2},
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shortestPathBinaryMatrix(tt.grid); got != tt.want {
				t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
