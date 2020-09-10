package question_1281_1290

import (
	"testing"
)

func Test_removeCoveredIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 4}, {3, 6}, {2, 8}}, 2},
		{[][]int{{1, 2}, {1, 4}, {3, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeCoveredIntervals(tt.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
