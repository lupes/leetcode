package question_451_460

import (
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {1, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMinArrowShots(tt.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
