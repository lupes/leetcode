package question_1261_1270

import (
	"testing"
)

func Test_countSquares(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, 15},
		{[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, 7},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countSquares(tt.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
