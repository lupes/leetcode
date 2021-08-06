package question_171_180

import (
	"testing"
)

func Test_calculateMinimumHP(t *testing.T) {
	tests := []struct {
		dungeon [][]int
		want    int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{[][]int{{0, 0}}, 1},
		{[][]int{{0, 0, 0}, {1, 1, -1}}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := calculateMinimumHP(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
