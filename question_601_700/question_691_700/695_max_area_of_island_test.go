package question_691_700

import (
	"testing"
)

func Test_maxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{nil, 0},
		{[][]int{{1}}, 1},
		{[][]int{{0}}, 0},
		{[][]int{{0, 1, 1}}, 2},
		{[][]int{{0, 1, 1, 1, 0}}, 3},
		{[][]int{{0, 1, 1, 1, 0}, {1, 0, 1, 0, 0}}, 4},
		{[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxAreaOfIsland(tt.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
