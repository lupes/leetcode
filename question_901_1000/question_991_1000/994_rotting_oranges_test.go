package question_991_1000

import (
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{2, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := orangesRotting(tt.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
