package question_431_440

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
		{[][]int{{1, 2}, {2, 4}, {2, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
