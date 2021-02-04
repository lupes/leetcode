package question_1441_1450

import (
	"testing"
)

func Test_minTime(t *testing.T) {
	tests := []struct {
		n        int
		edges    [][]int
		hasApple []bool
		want     int
	}{
		{1, [][]int{}, []bool{true}, 0},
		{2, [][]int{{0, 1}}, []bool{true, true}, 2},
		{4, [][]int{{0, 2}, {0, 3}, {1, 2}}, []bool{true, true, false, false}, 4},
		{7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{true, false, true, false, true, true, false}, 8},
		{7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, false, true, false}, 6},
		{7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, false, false, false, true, false}, 4},
		{7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, false, false, false, false, false}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minTime(tt.n, tt.edges, tt.hasApple); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
